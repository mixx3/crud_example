package app

import (
	"crud_example/src/api"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// Log godoc
// @Summary Status
// @Tags log
// @Accept */*
// @Produce json
// @Success 200
// @Router /v1/api/status [get]
func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "weight tracker API running smoothly",
		}

		c.JSON(http.StatusOK, response)
	}
}

// Log godoc
// @Summary Create Log
// @Tags log
// @Accept */*
// @Produce json
// @Success 200
// @Router /v1/api/log [post]
func (s *Server) CreateLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var newLog api.LogPostSchema
		err := c.ShouldBindJSON(&newLog)

		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		err = s.logService.Create(&newLog)

		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "new log created",
		}

		c.JSON(http.StatusOK, response)
	}
}

// HealthCheck godoc
// @Summary Get log by id
// @Tags log
// @Accept */*
// @Produce json
// @Success 200
// @Router /v1/api/log/{id} [get]
func (s *Server) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusUnprocessableEntity, nil)
			return
		}
		res, err := s.logService.Get(id)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

// HealthCheck godoc
// @Summary Get all
// @Tags log
// @Accept */*
// @Produce json
// @Success 200
// @Router /v1/api/log [get]
func (s *Server) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		res, err := s.logService.GetAll()
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusNotFound, nil)
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
