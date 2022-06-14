package handler

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	// UserService     model.UserService
}

type Config struct {
	R               *gin.Engine
	TimeoutDuration time.Duration
}

func NewHandler(config *Config) {

	h := &Handler{
		// UserService: config.UserService,
	}

	h.routes(config.R)

}
