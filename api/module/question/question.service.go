package question

import (
	"encoding/json"
	"net/http"
	"slambook/api/middlewere"
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

func (service *questionService) createQuestionHandler(ctx *gin.Context) {

	var questionDTO QuestionDTO
	if valid := validation.Bind(ctx, &questionDTO); !valid {
		return
	}

	if valid := validation.BindWith(ctx, &questionDTO); !valid {
		return
	}

	var questionParam QuestionParam
	if valid := validation.BindUri(ctx, &questionParam); !valid {
		return
	}

	var user middlewere.User
	reqUser := ctx.Request.Header.Get("user")
	json.Unmarshal([]byte(reqUser), &user)

	c := ctx.Request.Context()

	question, err := service.questionRepository.create(c, user, questionParam, questionDTO)

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
		Result:  question,
	})

}
func (service *questionService) readQuestionHandler(ctx *gin.Context) {

}
func (service *questionService) updateQuestionHandler(ctx *gin.Context) {

}
func (service *questionService) deleteQuestionHandler(ctx *gin.Context) {

}
