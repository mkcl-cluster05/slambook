package book

import (
	"slambook/api/middlewere"

	"github.com/gin-gonic/gin"
)

type BookRoute interface {
	Route(*gin.Engine)
}

type bookRoute struct {
	service BookService
}

func NewBookRoute(service BookService) BookRoute {
	return &bookRoute{
		service: service,
	}
}

func (handler *bookRoute) Route(router *gin.Engine) {
	book := router.Group("/book")
	{
		book.POST("", middlewere.Auth(), handler.service.createBookHandler)
		book.GET("/:bookId", middlewere.Auth(), handler.service.readBookHandler)
		book.GET("", middlewere.Auth(), handler.service.readAllBookHandler)
		book.PATCH("/:bookId", middlewere.Auth(), handler.service.updateBookHandler)
		book.DELETE("/:bookId", middlewere.Auth(), handler.service.deleteBookHandler)
	}
}
