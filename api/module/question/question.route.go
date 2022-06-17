package question

import (
	"slambook/api/middlewere"

	"github.com/gin-gonic/gin"
)

type QuestionRoute interface {
	Route(*gin.Engine)
}

type questionRoute struct {
	service QuestionService
}

func NewQuetionRoute(service QuestionService) QuestionRoute {
	return &questionRoute{
		service: service,
	}
}

func (handler *questionRoute) Route(router *gin.Engine) {

	question := router.Group("/book")
	{
		question.POST("/:bookId/question", middlewere.Auth(), handler.service.createQuestionHandler)
		question.GET("/:bookId/question", middlewere.Auth(), handler.service.readQuestionHandler)
		question.PATCH("/:bookId/question", middlewere.Auth(), handler.service.updateQuestionHandler)
		question.DELETE("/:bookId/question", middlewere.Auth(), handler.service.deleteQuestionHandler)

	}

}
