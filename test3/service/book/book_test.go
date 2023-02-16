package book

import (
	"gits/test3/models/book"
	"gits/test3/models/user"
	bookRepo "gits/test3/repository/book/mocks"
	userRepo "gits/test3/repository/user/mocks"
	"gits/test3/request"
	"testing"

	"github.com/stretchr/testify/assert"
)

var req = request.BookReq{
	Name:        "buku 1",
	PublisherId: 1,
}

var userId = float64(1)

func TestCreateBookUserNil(t *testing.T) {
	bookRepoMock := bookRepo.NewBookRepository(t)
	UserRepoMock := userRepo.NewUserRepository(t)
	BookService := NewBookService(bookRepoMock, UserRepoMock)

	UserRepoMock.Mock.On("GetById", userId).Return(nil).Once()
	create, err := BookService.CreateBook(req, userId)
	assert.Nil(t, create)
	assert.NotNil(t, err)
}

func TestCreateBook(t *testing.T) {
	bookRepoMock := bookRepo.NewBookRepository(t)
	userRepoMock := userRepo.NewUserRepository(t)
	BookService := NewBookService(bookRepoMock, userRepoMock)

	var userMock user.User
	bookMock := book.Book{
		Id:          0,
		Name:        "buku 1",
		PublisherId: 1,
	}

	userRepoMock.Mock.On("GetById", userId).Return(&userMock)
	bookRepoMock.Mock.On("Create", bookMock).Return(nil)
	create, err := BookService.CreateBook(req, userId)
	assert.Nil(t, err)
	assert.NotNil(t, create)
}
