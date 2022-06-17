package validation

import (
	"fmt"
	"net/http"
	r "slambook/utils/response"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func jkb(key string) string {
	if len(key) == 0 {
		return ""
	}
	return strings.ToLower(key[:1]) + key[1:]
}

func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", jkb(fe.Field()))
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", jkb(fe.Field()), fe.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s", jkb(fe.Field()), fe.Param())
	case "email":
		return "invalid email format"
	case "len":
		return fmt.Sprintf("%s must be %s characters long", jkb(fe.Field()), fe.Param())
	case "questionType":
		return fmt.Sprintf("%s must be of type %s ", jkb(fe.Field()), "MCQ,Descriptive,Image,Audio,Video,GIF")
	}
	return fmt.Sprintf("%s is not valid", jkb(fe.Field()))
}

func Bind(ctx *gin.Context, req interface{}) bool {

	if ctx.ContentType() != "application/json" {

		msg := fmt.Sprintf("request %s only accepts Content-Type application/json ", ctx.FullPath())

		ctx.JSON(http.StatusBadRequest, r.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request",
			Error:   msg,
		})

		return false
	}

	if err := ctx.ShouldBind(req); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {

			var invalidArgs []string
			for _, fe := range errs {
				invalidArgs = append(invalidArgs, getErrorMessage(fe))
			}

			ctx.JSON(http.StatusBadRequest, r.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Error:   invalidArgs,
			})

			return false

		}

		ctx.JSON(http.StatusInternalServerError, r.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Error:   err.Error(),
		})

		return false

	}
	return true

}

func BindUri(ctx *gin.Context, req interface{}) bool {

	if err := ctx.ShouldBindUri(req); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {

			var invalidArgs []string
			for _, fe := range errs {
				invalidArgs = append(invalidArgs, getErrorMessage(fe))
			}

			ctx.JSON(http.StatusBadRequest, r.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Error:   invalidArgs,
			})

			return false

		}

		ctx.JSON(http.StatusInternalServerError, r.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Error:   err.Error(),
		})

		return false

	}
	return true
}

func BindWith(ctx *gin.Context, req interface{}) bool {

	if err := ctx.ShouldBindWith(req, binding.Query); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {

			var invalidArgs []string
			for _, fe := range errs {
				invalidArgs = append(invalidArgs, getErrorMessage(fe))
			}

			ctx.JSON(http.StatusBadRequest, r.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Error:   invalidArgs,
			})

			return false

		}

		ctx.JSON(http.StatusInternalServerError, r.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Error:   err.Error(),
		})

		return false

	}
	return true
}
