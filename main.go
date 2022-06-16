package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"slambook/api"
	"slambook/utils/config"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := run(); err != nil {
		log.Fatalf("error : %v", err)
	}

}

func run() error {

	if err := setupEnvironment(); err != nil {
		return err
	}

	router, err := api.InitRouter()

	if err != nil {
		return err
	}

	server := http.Server{
		Addr:           fmt.Sprintf(":%s", config.ServerConfig.Port),
		Handler:        router,
		ReadTimeout:    config.ServerConfig.ReadTimeout * time.Second,
		WriteTimeout:   config.ServerConfig.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err = server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func setupEnvironment() error {

	env := flag.String("env", "dev", "To set environment dev/stg/prod")

	flag.Parse()

	if *env != "dev" && *env != "stg" && *env != "prod" {
		return fmt.Errorf("invalid environment type. check --help for to check env options")
	}

	if *env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := config.Setup(*env); err != nil {
		return err
	}

	return nil
}
