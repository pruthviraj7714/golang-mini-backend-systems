package api

import (
	"log"
	"net/http"
	"realtime-notification-service/internal/websocket"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	manager := websocket.NewWebsocketManager()

	r.GET("/ws", func(c *gin.Context) {
		userId := c.Query("userId")

		if userId == "" {
			log.Println("User ID is required")
			return
		}

		websocket.StartWebsocketServer(c, manager, userId)
	})

	r.GET("/connections", func(c *gin.Context) {
		websocket.GetConnections(c, manager)
	})

	r.POST("/notifications", func(c *gin.Context) {
		var req struct {
			UserId  string `json:"userId"`
			Message string `json:"message"`
		}

		if err := c.ShouldBindBodyWithJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"error":   "Invalid JSON",
			})
			return
		}

		err := manager.SendNotification(req.UserId, req.Message)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Message Successfully Sent to the user",
		})
	})

	r.Run()

}
