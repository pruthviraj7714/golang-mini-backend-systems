package websocket

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Validate request origin for security
	CheckOrigin: func(r *http.Request) bool {
		// In production, validate against allowed domains
		return true
	},
}

func StartWebsocketServer(c *gin.Context) {

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}

	defer ws.Close()

	log.Println("Client connected successfully!")

	for {
		messageType, messagePayload, err := ws.ReadMessage()

		if err != nil {
			log.Println("Error reading message or client disconnected:", err)
			break
		}

		log.Printf("Received message: %s\n", messagePayload)

		// Echo the same message back to the client
		err = ws.WriteMessage(messageType, messagePayload)
		if err != nil {
			log.Println("Error writing message:", err)
			break
		}

	}

}
