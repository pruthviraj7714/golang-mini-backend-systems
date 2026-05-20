package app

import (
	"banking-system/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {

	cfg := config.LoadConfig()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	// authRouter := router.Group("/auth")

	router.Run(":" + cfg.Port)

}
