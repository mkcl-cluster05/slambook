package auth

import (
	"net/http"
	"slambook/utils/binding"
	r "slambook/utils/response"

	"github.com/gin-gonic/gin"
)

const (
	USER_FOUND     = "user already present"
	USER_NOT_FOUND = "user not present"
)

type AuthService interface {
	registerHandler(*gin.Context)
	loginHandler(*gin.Context)
	changePasswordHandler(*gin.Context)
	forgotPasswordHandler(*gin.Context)
}

type authService struct {
	authRepository AuthRepository
}

func NewAuthService(authRepository AuthRepository) AuthService {
	return &authService{
		authRepository: authRepository,
	}
}

func (service *authService) registerHandler(ctx *gin.Context) {

	var registerDTO RegisterDTO
	if valid := binding.BindData(ctx, &registerDTO); !valid {
		return
	}

	isUserPresent := service.authRepository.checkUser(ctx, registerDTO.Email)

	if isUserPresent {
		ctx.JSON(http.StatusBadRequest, r.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Error:   USER_FOUND,
		})
		return
	}

	c := ctx.Request.Context()
	auth, err := service.authRepository.register(c, registerDTO)
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
		Result:  auth,
	})
}

func (service *authService) loginHandler(ctx *gin.Context) {

	var loginDTO LoginDTO
	if valid := binding.BindData(ctx, &loginDTO); !valid {
		return
	}

	isUserPresent := service.authRepository.checkUser(ctx, loginDTO.Email)

	if !isUserPresent {
		ctx.JSON(http.StatusNotFound, r.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "error",
			Error:   USER_NOT_FOUND,
		})
		return
	}

	c := ctx.Request.Context()
	auth, err := service.authRepository.login(c, loginDTO)
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
		Result:  auth,
	})

}
func (service *authService) changePasswordHandler(ctx *gin.Context) {

}
func (service *authService) forgotPasswordHandler(ctx *gin.Context) {

}
