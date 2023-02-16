package book

import (
	"errors"
	"gits/test3/models/book"
	bookPort "gits/test3/port/book"
	userPort "gits/test3/port/user"
	"gits/test3/request"
)

type BookService struct {
	bookRepo bookPort.BookRepository
	userRepo userPort.UserRepository
}

func NewBookService(
	bookRepo bookPort.BookRepository,
	userRepo userPort.UserRepository,
) BookService {
	return BookService{
		bookRepo: bookRepo,
		userRepo: userRepo,
	}
}

func (s *BookService) CreateBook(req request.BookReq, userInfo interface{}) (*book.Book, error) {
	user, err := s.userRepo.GetById(userInfo.(float64))
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user login 404")
	}

	book := book.Book{
		Name:        req.Name,
		AuthorId:    user.AuthorId,
		PublisherId: req.PublisherId,
	}

	err = s.bookRepo.Create(book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *BookService) ListBook(userInfo interface{}) ([]book.Book, error) {
	user, err := s.userRepo.GetById(userInfo.(float64))
	if err != nil {
		return nil, err
	}

	books := s.bookRepo.List(user.AuthorId)
	return books, nil
}

func (s *BookService) UpdateBook(bookId string, req request.BookReq, userInfo interface{}) (*book.Book, error) {
	user, err := s.userRepo.GetById(userInfo.(float64))
	if err != nil {
		return nil, err
	}

	dataBook := s.bookRepo.GetById(bookId)
	dataBook.Name = req.Name
	dataBook.PublisherId = req.PublisherId
	dataBook.AuthorId = user.AuthorId

	err = s.bookRepo.Update(dataBook)
	if err != nil {
		return nil, err
	}

	return &dataBook, nil
}

func (s *BookService) DeleteBook(bookId string) error {
	dataBook := s.bookRepo.GetById(bookId)
	err := s.bookRepo.Delete(dataBook)
	if err != nil {
		return err
	}

	return nil
}
