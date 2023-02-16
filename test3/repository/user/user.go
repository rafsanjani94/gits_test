package user

import (
	"context"
	"encoding/json"
	"gits/test3/config"
	"gits/test3/conv"
	"gits/test3/models/user"
	userPort "gits/test3/port/user"
	"time"
)

type UserRepository struct {
}

func NewUserRepository() userPort.UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Login(username, password string) user.User {
	var userLogin user.User
	config.DB.Debug().Where(&user.User{
		Username: username,
		Password: string(password),
	}).First(&userLogin)

	return userLogin
}

func (r *UserRepository) GetById(userId float64) (*user.User, error) {
	redisUser, err := config.RedisClient.Get(context.TODO(), "user:"+conv.Float64ToString(userId)).Result()
	if err != nil {
		return nil, err
	}

	var user *user.User
	err = json.Unmarshal([]byte(redisUser), &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) SetLogin(userId string, loginData []byte, exp time.Duration) error {
	_, err := config.RedisClient.Set(context.TODO(), "user:"+userId, string(loginData), exp).Result()
	if err != nil {
		return err
	}

	return nil
}
