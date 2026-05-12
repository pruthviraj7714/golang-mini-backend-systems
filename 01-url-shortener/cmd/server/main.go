package main

import (
	"net/http"
	"url-shortener/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	r.Run(":" + config.LoadConfig().Port)
}
