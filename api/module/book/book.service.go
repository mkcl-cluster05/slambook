package book

import (
	"encoding/json"
	"net/http"
	"slambook/api/middlewere"
	r "slambook/utils/response"
	"slambook/utils/validation"

	"github.com/gin-gonic/gin"
)

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

func (service *bookService) createBookHandler(ctx *gin.Context) {

	var bookDTO BookDTO
	if valid := validation.BindData(ctx, &bookDTO); !valid {
		return
	}

	var user middlewere.User
	reqUser := ctx.Request.Header.Get("user")
	json.Unmarshal([]byte(reqUser), &user)

	c := ctx.Request.Context()

	book, err := service.bookRepository.create(c, user, bookDTO)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, r.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, r.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Result:  book,
	})
}
func (service *bookService) readAllBookHandler(ctx *gin.Context) {

	var user middlewere.User
	reqUser := ctx.Request.Header.Get("user")
	json.Unmarshal([]byte(reqUser), &user)

	c := ctx.Request.Context()

	books, err := service.bookRepository.readAll(c, user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, r.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, r.SuccessResponse{
		Status:  http.StatusOK,
		Message: "success",
		Result:  books,
	})

}
func (service *bookService) readBookHandler(ctx *gin.Context) {

	var bookParam BookParam
	if valid := validation.BindUri(ctx, &bookParam); !valid {
		return
	}

	var user middlewere.User
	reqUser := ctx.Request.Header.Get("user")
	json.Unmarshal([]byte(reqUser), &user)

	c := ctx.Request.Context()

	book, err := service.bookRepository.read(c, user, bookParam.BookId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, r.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, r.SuccessResponse{
		Status:  http.StatusOK,
		Message: "success",
		Result:  book,
	})

}
func (service *bookService) updateBookHandler(ctx *gin.Context) {

	var bookParam BookParam
	if valid := validation.BindUri(ctx, &bookParam); !valid {
		return
	}

	var bookDTO BookDTO
	if valid := validation.BindData(ctx, &bookDTO); !valid {
		return
	}

	var user middlewere.User
	reqUser := ctx.Request.Header.Get("user")
	json.Unmarshal([]byte(reqUser), &user)

	c := ctx.Request.Context()

	book, err := service.bookRepository.update(c, user, bookParam.BookId, bookDTO)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, r.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, r.SuccessResponse{
		Status:  http.StatusOK,
		Message: "success",
		Result:  book,
	})

}
func (service *bookService) deleteBookHandler(ctx *gin.Context) {

	var bookParam BookParam
	if valid := validation.BindUri(ctx, &bookParam); !valid {
		return
	}

	var user middlewere.User
	reqUser := ctx.Request.Header.Get("user")
	json.Unmarshal([]byte(reqUser), &user)

	c := ctx.Request.Context()

	err := service.bookRepository.delete(c, user, bookParam.BookId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, r.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, r.SuccessResponse{
		Status:  http.StatusOK,
		Message: "success",
		Result:  "deleted successfully",
	})

}
