package handler

import (
	"mathalama/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
	systemSvc service.SystemService
}

func NewAppHandler(systemSvc service.SystemService) *AppHandler {
	return &AppHandler{systemSvc: systemSvc}
}

func (h *AppHandler) HealthCheck(c *gin.Context) {
	if err := h.systemSvc.CheckHealth(c.Request.Context()); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "error",
			"message": "System is unhealthy",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "System is running smoothly",
	})
}
