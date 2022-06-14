package api

import (
	"slambook/api/handler"
	"slambook/datasource"
	"time"

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

	return router, nil
}

func setupHandler(router *gin.Engine, ds *datasource.DataSource) {

	handler.NewHandler(&handler.Config{
		R: router,
		// UserService : getUserService(ds),
		TimeoutDuration: 5 * time.Second,
	})
}
