package book

import "gits/test3/models/book"

type BookRepository interface {
	Create(book book.Book) error
	List(authorId int) []book.Book
	GetById(id string) book.Book
	Update(book book.Book) error
	Delete(book book.Book) error
	ListAll() []book.Book
}
