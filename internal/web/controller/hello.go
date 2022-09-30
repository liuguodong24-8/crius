package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello ...
func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "hello world",
	})
}
