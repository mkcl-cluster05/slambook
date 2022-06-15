package binding

import (
	"fmt"
	"net/http"
	r "slambook/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type InvalidArgument struct {
	Field string
	Value interface{}
	Tag   string
	Param string
}

func BindData(ctx *gin.Context, req interface{}) bool {

	if ctx.ContentType() != "application/json" {

		msg := fmt.Sprintf("%s only accepts Content-Type application/json ", ctx.FullPath())

		ctx.JSON(http.StatusBadRequest, r.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request",
			Error:   msg,
		})

		return false
	}

	if err := ctx.ShouldBind(req); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {

			var invalidArgs []InvalidArgument

			for _, er := range errs {

				invalidArgs = append(invalidArgs, InvalidArgument{
					er.Field(),
					er.Value(),
					er.Tag(),
					er.Param(),
				})

			}

			ctx.JSON(http.StatusBadRequest, r.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "invalid request",
				Error:   invalidArgs,
			})

			return false

		}

		ctx.JSON(http.StatusInternalServerError, r.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
			Error:   "internal server error",
		})

		return false

	}

	return true

}
