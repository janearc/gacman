package main

import (
	"gacman/models"
	"github.com/gorilla/websocket"
	"time"

	"log"
)

const (
	wsEndpoint = "ws://localhost:8080" // Replace with your WebSocket server's address
)

// StartDaemon sets up the WebSocket connection and sends data periodically.
func StartDaemon() {
	// Establish WebSocket connection
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer conn.Close()

	log.Println("Connected to WebSocket server.")

	// Example object to send
	obj := models.NewObject(1, 2, 3, "grass", 5.0, "tree")

	// Periodically send data
	ticker := time.NewTicker(5 * time.Second) // Send data every 5 seconds
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Serialize object to JSON
			jsonString := obj.ToJSON()

			// Send the JSON data to the server
			err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString))
			if err != nil {
				log.Printf("Error sending message: %v", err)
				return
			}
			log.Printf("Sent data: %s", jsonString)
		}
	}
}

func main() {
	StartDaemon()
}
