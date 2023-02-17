package app

import (
	"crud_example/src/api"
	"github.com/gin-gonic/gin"

	"log"
)

type Server struct {
	router     *gin.Engine
	logService api.LogService
}

func NewServer(router *gin.Engine, logService api.LogService) *Server {
	return &Server{
		router:     router,
		logService: logService,
	}
}

func (s *Server) Run() error {
	// run function that initializes the routes
	r := s.Routes()

	// run the server through the router
	err := r.Run()

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
