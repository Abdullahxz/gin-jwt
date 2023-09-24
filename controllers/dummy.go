package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Dummy controller, could be anything here.

func DummyController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "This is working.",
	})
}
