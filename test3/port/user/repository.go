package user

import (
	"gits/test3/models/user"
	"time"
)

type UserRepository interface {
	Login(username, password string) user.User
	GetById(userId float64) (*user.User, error)
	SetLogin(userId string, loginData []byte, exp time.Duration) error
}
