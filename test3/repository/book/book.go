package book

import (
	"gits/test3/config"
	"gits/test3/models/book"
	bookPort "gits/test3/port/book"
)

type BookRepository struct {
}

func NewBookRepository() bookPort.BookRepository {
	return &BookRepository{}
}

func (r *BookRepository) Create(book book.Book) error {
	return config.DB.Debug().Create(&book).Error
}

func (r *BookRepository) List(authorId int) []book.Book {
	var books []book.Book
	config.DB.Debug().Where("author_id", authorId).Find(&books)
	return books
}

func (r *BookRepository) GetById(id string) book.Book {
	var dataBook book.Book
	config.DB.Debug().Where("id", id).First(&dataBook)
	return dataBook
}

func (r *BookRepository) Update(book book.Book) error {
	return config.DB.Debug().Save(&book).Error
}

func (r *BookRepository) Delete(book book.Book) error {
	return config.DB.Debug().Delete(&book).Error
}

func (r *BookRepository) ListAll() []book.Book {
	var books []book.Book
	config.DB.Debug().Find(&books)
	return books
}
