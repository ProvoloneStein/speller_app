package handler

import (
	"Nexign/internal/service"
	"github.com/gin-gonic/gin"
	"os"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/text", h.postText)
		v1.POST("/texts", h.postMany)
	}

	return router
}
