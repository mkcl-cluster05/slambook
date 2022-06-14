package auth

import "github.com/gin-gonic/gin"

type AuthHandler struct {
	service AuthService
}

func (handler *AuthHandler) AuthHandler(router *gin.Engine) {

	auth := router.Group("/auth")
	{
		auth.POST("/register", handler.service.registerHandler)
		auth.POST("/login", handler.service.loginHandler)
		auth.POST("/changePassword", handler.service.changePasswordHandler)
		auth.POST("/forgotPassword", handler.service.forgotPasswordHandler)
	}

}
