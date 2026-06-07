package websocket

import (
	"errors"
	"log"
	"net/http"
	"realtime-notification-service/internal/models"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
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
	db      *gorm.DB
	mutex   sync.RWMutex
}

func NewWebsocketManager(db *gorm.DB) *WebsocketManager {
	return &WebsocketManager{
		clients: make(map[string]*websocket.Conn),
		db:      db,
		mutex:   sync.RWMutex{},
	}
}

func (m *WebsocketManager) RegisterClient(conn *websocket.Conn, userId string) {
	defer m.mutex.Unlock()
	m.mutex.Lock()

	m.clients[userId] = conn
	log.Printf("Client connected: %v, User ID: %s", conn.RemoteAddr(), userId)
}

func (m *WebsocketManager) UnregisterClient(userId string) {
	defer m.mutex.Unlock()
	m.mutex.Lock()

	delete(m.clients, userId)
	log.Printf("Client disconnected: %s", userId)
}

func (m *WebsocketManager) BroadcastMessage(message []byte) {
	defer m.mutex.Unlock()
	m.mutex.Lock()
	for _, conn := range m.clients {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Error broadcasting message:", err)
			continue
		}
	}
}

func (m *WebsocketManager) SendNotification(userId string, message string) error {
	conn, exists := m.clients[userId]

	if !exists {
		return errors.New("user not connected")
	}

	err := m.db.Create(&models.Notification{
		UserID:  userId,
		Message: message,
	}).Error

	if err != nil {
		return err
	}

	err = conn.WriteMessage(websocket.TextMessage, []byte(message))

	if err != nil {
		return err
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

func GetConnections(manager *WebsocketManager) ([]string, error) {
	defer manager.mutex.RUnlock()

	manager.mutex.RLock()

	var connections []string

	for userId := range manager.clients {
		connections = append(connections, userId)
	}

	return connections, nil

}

func GetNotificationsHistory(manager *WebsocketManager, userId string) ([]models.Notification, error) {

	var notifications []models.Notification

	err := manager.db.
		Where("user_id = ?", userId).
		Order("created_at DESC").
		Limit(20).
		Find(&notifications).
		Error

	if err != nil {
		return []models.Notification{}, err
	}

	return notifications, nil
}
