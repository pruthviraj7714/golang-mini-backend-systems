package main

import (
	"api-gateway/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.LoggingMiddleware())

	r.GET("/health", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "gateway is running",
		})
	})

	r.Run(":8080")
}
