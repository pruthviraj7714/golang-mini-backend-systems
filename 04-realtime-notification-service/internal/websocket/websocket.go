package websocket

import (
	"fmt"
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

type WebsocketManager struct {
	clients map[string]*websocket.Conn
}

func NewWebsocketManager() *WebsocketManager {
	return &WebsocketManager{
		clients: make(map[string]*websocket.Conn),
	}
}

func (m *WebsocketManager) RegisterClient(conn *websocket.Conn, userId string) {
	m.clients[userId] = conn
	log.Printf("Client connected: %v, User ID: %s", conn.RemoteAddr(), userId)
}

func (m *WebsocketManager) UnregisterClient(userId string) {
	delete(m.clients, userId)
	log.Printf("Client disconnected: %s", userId)
}

func (m *WebsocketManager) BroadcastMessage(message []byte) {
	for _, conn := range m.clients {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Error broadcasting message:", err)
			continue
		}
	}
}

func (m *WebsocketManager) SendNotification(userId, message string) error {
	for uId, conn := range m.clients {
		if uId == userId {
			err := conn.WriteMessage(websocket.TextMessage, []byte(message))
			return err
		}
	}

	return nil
}

func StartWebsocketServer(c *gin.Context, manager *WebsocketManager, userId string) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}

	manager.RegisterClient(ws, userId)

	fmt.Println("active clients:", len(manager.clients))

	defer manager.UnregisterClient(userId)

	defer func() {
		ws.Close()
		log.Printf("Client disconnected: %v", ws.RemoteAddr())
	}()

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

func GetConnections(c *gin.Context, manager *WebsocketManager) {
	var connections []string

	for userId, _ := range manager.clients {
		connections = append(connections, userId)
	}

	c.JSON(http.StatusOK, gin.H{
		"connectedClients": connections,
	})

}
