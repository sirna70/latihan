package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	__USERNAME__ = "napit"
	__PASSWORD__ = "napit123"
)

func Auth(c *gin.Context) bool {
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Auhtorization Required"})
		return false
	}
	if username != __USERNAME__ || password != __PASSWORD__ {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return false
	}
	return true
}
