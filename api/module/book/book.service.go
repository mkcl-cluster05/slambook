package book

import "github.com/gin-gonic/gin"

type BookService interface {
	createBookHandler(*gin.Context)
	readAllBookHandler(*gin.Context)
	readBookHandler(*gin.Context)
	updateBookHandler(*gin.Context)
	deleteBookHandler(*gin.Context)
}

type bookService struct {
	bookRepository BookRepository
}

func NewBookService(bookRepository BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepository,
	}
}

func (service *bookService) createBookHandler(ctx *gin.Context)  {}
func (service *bookService) readAllBookHandler(ctx *gin.Context) {}
func (service *bookService) readBookHandler(ctx *gin.Context)    {}
func (service *bookService) updateBookHandler(ctx *gin.Context)  {}
func (service *bookService) deleteBookHandler(ctx *gin.Context)  {}
