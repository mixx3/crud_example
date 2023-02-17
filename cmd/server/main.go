package main

import (
	api2 "crud_example/pkg/api"
	"crud_example/pkg/app"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	router := gin.Default()
	router.Use(cors.Default())
	logService := api2.NewLogService(api2.NewFakeRepo())
	server := app.NewServer(router, logService)
	err := server.Run()
	if err != nil {
		return err
	}
	return nil
}