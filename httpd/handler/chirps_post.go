package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChirpsPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// use the gin.Context to respond to the incoming request
		c.JSON(http.StatusOK, map[string]string{
			"hello": "Found me",
		})
	}
}
