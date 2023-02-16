package handler

import (
	bookRepo "gits/test3/repository/book"
	userRepo "gits/test3/repository/user"
	"gits/test3/request"
	bookService "gits/test3/service/book"
	loginService "gits/test3/service/login"
	"net/http"

	"github.com/gin-gonic/gin"
)

var BookRepo = bookRepo.NewBookRepository()
var UserRepo = userRepo.NewUserRepository()
var BookService = bookService.NewBookService(BookRepo, UserRepo)

func Login(c *gin.Context) {
	var req request.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	login, token, err := loginService.Login(req)
	if err != nil {
		if err.Error() == "401" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "invalid credential",
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"token":  token,
		"data":   login,
	})
}

func CreateBook(c *gin.Context) {
	var req request.BookReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userInfo, _ := c.Get("userInfo")
	create, err := BookService.CreateBook(req, userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data":   create,
	})
}

func ListBook(c *gin.Context) {
	userInfo, _ := c.Get("userInfo")
	list, err := BookService.ListBook(userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   list,
	})
}

func UpdateBook(c *gin.Context) {
	var req request.BookReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userInfo, _ := c.Get("userInfo")
	create, err := BookService.UpdateBook(c.Param("id"), req, userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   create,
	})
}

func DeleteBook(c *gin.Context) {
	err := BookService.DeleteBook(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}
