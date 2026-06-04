package api

import (
	"realtime-notification-service/internal/websocket"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	r.Get("/ws", websocket.StartWebsocketServer)

	r.Run()

}
