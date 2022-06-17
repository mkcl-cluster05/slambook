package question

import (
	"net/http"
	r "slambook/utils/response"
	"slambook/utils/validation"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("questionType", QuestionTypeValidation)
	}

	return &questionService{
		questionRepository: questionRepository,
	}
}

func (repo *questionService) createQuestionHandler(ctx *gin.Context) {

	var questionDTO QuestionDTO
	if valid := validation.Bind(ctx, &questionDTO); !valid {
		return
	}

	if valid := validation.BindWith(ctx, &questionDTO); !valid {
		return
	}

	ctx.JSON(http.StatusOK, r.SuccessResponse{
		Status:  http.StatusOK,
		Message: "success",
		Result:  "ok",
	})

}
func (repo *questionService) readQuestionHandler(ctx *gin.Context) {

}
func (repo *questionService) updateQuestionHandler(ctx *gin.Context) {

}
func (repo *questionService) deleteQuestionHandler(ctx *gin.Context) {

}
