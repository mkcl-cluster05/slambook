package book

import "github.com/gin-gonic/gin"

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
		book.POST("/", handler.service.createBookHandler)
		book.GET("/", handler.service.readAllBookHandler)
		book.GET("/:bookId", handler.service.readBookHandler)
		book.PATCH("/:bookId", handler.service.updateBookHandler)
		book.DELETE("/:bookId", handler.service.deleteBookHandler)
	}
}
