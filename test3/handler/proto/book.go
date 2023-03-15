package proto

import (
	pb "gits/test3/proto/book"
	bookRepo "gits/test3/repository/book"
	userRepo "gits/test3/repository/user"
	bookService "gits/test3/service/book"
)

var BookRepo = bookRepo.NewBookRepository()
var UserRepo = userRepo.NewUserRepository()
var BookService = bookService.NewBookService(BookRepo, UserRepo)

type BookServer struct {
	pb.UnimplementedBookServiceServer
	books []*pb.Book
}

func (s *BookServer) ListBook(req *pb.BooksQuery, stream pb.BookService_ListBookServer) error {
	books, err := BookService.ListBookAll()
	if err != nil {
		return err
	}

	for _, book := range books {
		bookStream := &pb.Book{
			Id:   int32(book.Id),
			Name: book.Name,
		}

		if err := stream.Send(bookStream); err != nil {
			return err
		}
	}

	return nil
}
