package auth

import (
	"net/http"
	"slambook/utils/binding"
	r "slambook/utils/response"

	"github.com/gin-gonic/gin"
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

	c := ctx.Request.Context()

	auth, err := service.authRepository.register(c, registerDTO)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, r.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, r.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data: AuthResponse{
			AccessToken: auth.AuthId,
		},
	})
}

func (service *authService) loginHandler(ctx *gin.Context) {

}
func (service *authService) changePasswordHandler(ctx *gin.Context) {

}
func (service *authService) forgotPasswordHandler(ctx *gin.Context) {

}
