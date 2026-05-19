package app

import (
	"log"
	"net/http"
	"worker-pool/internal/config"
	"worker-pool/internal/db"
	"worker-pool/internal/jobs"
	"worker-pool/internal/queue"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	cfg := config.LoadConfig()
	database := db.Connect(cfg.DB_URL)

	database.AutoMigrate(&jobs.Job{})

	queue := queue.NewJobQueue(100)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	r.POST("/jobs", func(c *gin.Context) {
		var CreateJobRequest struct {
			Type    string `json:"type"`
			Payload any    `json:"payload"`
		}

		if err := c.ShouldBindJSON(&CreateJobRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		job := jobs.Job{
			Type:    CreateJobRequest.Type,
			Status:  "PENDING",
			Payload: CreateJobRequest.Payload,
		}

		resp := database.Model(&jobs.Job{}).Create(&job)

		if resp.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
				"error":   resp.Error.Error(),
			})
			return
		}

		queue.Enqueue(&job)

		c.JSON(http.StatusOK, gin.H{
			"message": "Job successfully Pushed into queue",
		})
	})

	err := r.Run(":" + cfg.PORT)

	if err != nil {
		log.Fatal("API Server failed: ", err)
	}
}
