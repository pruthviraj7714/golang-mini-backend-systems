package app

import (
	"log"
	"net/http"
	"realtime-notification-service/internal/config"

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

func handleConnections(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("Error upgrading connection: %v", err)
		return
	}

	defer conn.Close()

	log.Println("Client connected successfully")

	for {

		messageType, message, err := conn.ReadMessage()

		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		log.Printf("Received message: %s", message)

		// Echo the received message back to the client
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Printf("Write error: %v", err)
			break
		}
	}

}

func Start() {
	cfg := config.New()

	http.HandleFunc("/ws", handleConnections)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
