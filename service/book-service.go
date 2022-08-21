package service

import (
	"fmt"
	"log"

	"github.com/AbdulQoheng/go-mvc/app/models"
	"github.com/AbdulQoheng/go-mvc/database/dto"
	"github.com/AbdulQoheng/go-mvc/database/entity"
	"github.com/mashingan/smapping"
)

// BookService is a ....
type BookService interface {
	Insert(b dto.BookCreateDTO) entity.Book
	Update(b dto.BookUpdateDTO) entity.Book
	Delete(b entity.Book)
	All() []entity.Book
	FindByID(bookID uint64) entity.Book
	IsAllowedToEdit(userID string, bookID uint64) bool
}

type bookService struct {
	bookModels models.BookModels
}

// NewBookService .....
func NewBookService(bookRepo models.BookModels) BookService {
	return &bookService{
		bookModels: bookRepo,
	}
}

func (service *bookService) Insert(b dto.BookCreateDTO) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.bookModels.InsertBook(book)
	return res
}

func (service *bookService) Update(b dto.BookUpdateDTO) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.bookModels.UpdateBook(book)
	return res
}

func (service *bookService) Delete(b entity.Book) {
	service.bookModels.DeleteBook(b)
}

func (service *bookService) All() []entity.Book {
	return service.bookModels.AllBook()
}

func (service *bookService) FindByID(bookID uint64) entity.Book {
	return service.bookModels.FindBookByID(bookID)
}

func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
	b := service.bookModels.FindBookByID(bookID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
