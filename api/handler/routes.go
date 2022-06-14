package handler

import "github.com/gin-gonic/gin"

func (h *Handler) routes(router *gin.Engine) {

	o := router.Group("/o")
	r := router.Group("/r")

	// r.Use() // use auth middlewere

	o.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"open": "seccess",
		})
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"restricted": "seccess",
		})
	})

}
