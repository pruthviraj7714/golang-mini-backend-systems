package app

import (
	"log"
	"net/http"
	"worker-pool/internal/config"
	"worker-pool/internal/db"
	"worker-pool/internal/jobs"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	cfg := config.LoadConfig()
	database := db.Connect(cfg.DB_URL)

	database.AutoMigrate(&jobs.Job{})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	err := r.Run(":" + cfg.PORT)

	if err != nil {
		log.Fatal("API Server failed: ", err)
	}
}
