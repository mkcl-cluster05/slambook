package auth

import "github.com/gin-gonic/gin"

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
}

func (service *authService) loginHandler(ctx *gin.Context) {

}
func (service *authService) changePasswordHandler(ctx *gin.Context) {

}
func (service *authService) forgotPasswordHandler(ctx *gin.Context) {

}
