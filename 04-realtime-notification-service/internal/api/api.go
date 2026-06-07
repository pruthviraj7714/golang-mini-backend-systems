package api

import (
	"log"
	"net/http"
	"realtime-notification-service/internal/config"
	"realtime-notification-service/internal/db"
	"realtime-notification-service/internal/websocket"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	cfg := config.New()

	database := db.Connect(cfg.DBURL)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	manager := websocket.NewWebsocketManager(database)

	r.GET("/ws", func(c *gin.Context) {
		userId := c.Query("userId")

		if userId == "" {
			log.Println("User ID is required")
			return
		}

		websocket.StartWebsocketServer(c, manager, userId)
	})

	r.GET("/connections", func(c *gin.Context) {
		connections, err := websocket.GetConnections(manager)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"connectedClients": connections,
		})
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

	r.GET("/notifications/:userId", func(c *gin.Context) {
		userId, exists := c.Params.Get("userId")

		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "User Id not found",
			})
			return
		}

		data, err := websocket.GetNotificationsHistory(manager, userId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"notifications": data,
		})
	})

	r.Run(":" + cfg.Port)

}
