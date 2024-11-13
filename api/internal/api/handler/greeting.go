package handler

import (
	"net/http"

	"github.com/circuit-shell/playlist-builder-back/pkg/greeting"
	"github.com/gin-gonic/gin"
)

// GreetingHandler handles greeting-related requests
type GreetingHandler struct{}

// NewGreetingHandler creates a new GreetingHandler
func NewGreetingHandler() *GreetingHandler {
	return &GreetingHandler{}
}

// GetGreeting handles GET requests for greetings
func (h *GreetingHandler) GetGreeting(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "name parameter is required",
		})
		return
	}

	message := greeting.GetGreeting(name)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
