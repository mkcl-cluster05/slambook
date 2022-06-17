package question

import (
	"github.com/gin-gonic/gin"
)

type QuestionService interface {
	createQuestionHandler(*gin.Context)
	readQuestionHandler(*gin.Context)
	updateQuestionHandler(*gin.Context)
	deleteQuestionHandler(*gin.Context)
}

type questionService struct {
	questionRepository QuestionRepository
}

func NewQuestionService(questionRepository QuestionRepository) QuestionService {
	return &questionService{
		questionRepository: questionRepository,
	}
}

func (repo *questionService) createQuestionHandler(ctx *gin.Context) {

}
func (repo *questionService) readQuestionHandler(ctx *gin.Context) {

}
func (repo *questionService) updateQuestionHandler(ctx *gin.Context) {

}
func (repo *questionService) deleteQuestionHandler(ctx *gin.Context) {

}
