package app

import (
	"crud_example/pkg/api"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// ApiStatus Log godoc
// @Summary Status
// @Tags log
// @Accept */*
// @Produce json
// @Success 200
// @Router /v1/status [get]
func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
		}

		c.JSON(http.StatusOK, response)
	}
}

// CreateLog Log godoc
// @Summary Create Log
// @Tags log
// @Accept json
// @Produce json
// @Param log body api.LogPostSchema true "schema"
// @Success 200
// @Router /v1/log [post]
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

// GetById HealthCheck godoc
// @Summary Get log by id
// @Tags log
// @Accept */*
// @Produce json
// @Success 200
// @Param        id    path     int  0  "id of db log"
// @Router /v1/log/{id} [get]
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

// GetAll HealthCheck godoc
// @Summary Get all
// @Tags log
// @Accept */*
// @Produce json
// @Success 200
// @Router /v1/log [get]
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
