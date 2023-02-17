package app

import (
	docs "crud_example/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) Routes() *gin.Engine {
	router := s.router
	docs.SwaggerInfo.Version = "0.2.1"
	// group all routes under /v1/api
	v1 := router.Group("/v1")
	{
		v1.GET("/status", s.ApiStatus())
		// prefix the user routes
		log := v1.Group("/log")
		log.POST("", s.CreateLog())
		log.GET("/:id", s.GetById())
		log.GET("", s.GetAll())
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
