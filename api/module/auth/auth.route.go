package auth

import (
	"slambook/api/middlewere"

	"github.com/gin-gonic/gin"
)

type AuthRoute interface {
	Route(*gin.Engine)
}
type authRoute struct {
	service AuthService
}

func NewAuthRoute(service AuthService) AuthRoute {
	return &authRoute{
		service: service,
	}
}

func (handler *authRoute) Route(router *gin.Engine) {

	auth := router.Group("/auth")
	{
		auth.POST("/register", handler.service.registerHandler)
		auth.POST("/login", handler.service.loginHandler)
		auth.POST("/change-password", middlewere.Auth(), handler.service.changePasswordHandler)
		auth.POST("/forgot-password", handler.service.forgotPasswordHandler)
	}
}
