package api

import (
	"net/http"
	"slambook/api/module/auth"
	"slambook/api/module/book"
	"slambook/datasource"
	r "slambook/utils/response"

	"github.com/gin-gonic/gin"
)

func InitRouter() (*gin.Engine, error) {

	ds, err := datasource.InitDS()

	if err != nil {
		return nil, err
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	setupHandler(router, ds)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, r.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "route not defined",
			Error:   "router not defined",
		})
	})
	return router, nil
}

func setupHandler(router *gin.Engine, ds *datasource.DataSource) {

	initAuthModule(router, ds)
	initBookModule(router, ds)

}

func initAuthModule(router *gin.Engine, ds *datasource.DataSource) {
	authRepository := auth.NewAuthRepository(ds.MongoDB)
	authService := auth.NewAuthService(authRepository)
	authRoute := auth.NewAuthRoute(authService)
	authRoute.Route(router)
}

func initBookModule(router *gin.Engine, ds *datasource.DataSource) {
	bookRepository := book.NewBookRepository(ds.MongoDB)
	bookService := book.NewBookService(bookRepository)
	bookRoute := book.NewBookRoute(bookService)
	bookRoute.Route(router)
}
