package login

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"gits/test3/conv"
	"gits/test3/models/user"
	"gits/test3/request"
	"net/http"
	"strings"
	"time"

	userRepo "gits/test3/repository/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var APPLICATION_NAME = "Books"
var LOGIN_EXPIRATION_DURATION = time.Duration(24) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("secrets")

var UserRepo = userRepo.NewUserRepository()

func GetMD5HashWithSum(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

type MyClaims struct {
	jwt.StandardClaims
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func Login(req request.LoginReq) (user.User, string, error) {
	var passInput = GetMD5HashWithSum(req.Password)
	var userLogin = UserRepo.Login(req.Username, passInput)
	if userLogin.Id == 0 {
		return userLogin, "", errors.New("401")
	}

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Id:       userLogin.Id,
		Username: userLogin.Username,
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return userLogin, "", err
	}

	loginData, err := json.Marshal(userLogin)
	if err != nil {
		return userLogin, "", err
	}

	UserRepo.SetLogin(conv.IntToString(userLogin.Id), loginData, LOGIN_EXPIRATION_DURATION)
	return userLogin, signedToken, nil
}

func Auth(c *gin.Context) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	//pencocokan token jwt
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("Signing method invalid")
		}
		return JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid token/not authorized",
		})

		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid token/not authorized",
		})

		c.Abort()
		return
	}

	c.Set("userInfo", claims["id"])
}
