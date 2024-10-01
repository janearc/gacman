package main

import (
	"net/http"

	"gacman/models" // Updated import path

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// handleConnections upgrades HTTP connections to WebSocket connections and processes messages.
func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("Failed to upgrade to WebSocket: %v", err)
		return
	}
	defer func() {
		if err := ws.Close(); err != nil {
			log.Errorf("Failed to close WebSocket connection: %v", err)
		}
	}()

	log.Info("Client connected.")

	// Create an example object to send
	obj := models.NewObject(1, 2, 3, "grass", 5.0, "tree")

	// Main loop: Send the object to the client periodically
	for {
		// Serialize the object to JSON
		jsonString := obj.ToJSON()

		// Send JSON data to the client
		err = ws.WriteMessage(websocket.TextMessage, []byte(jsonString))
		if err != nil {
			log.Errorf("Error sending message: %v", err)
			break
		}

		log.Infof("Sent data: %s", jsonString)
	}
}

func main() {
	// Set up logrus logging
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(log.InfoLevel)

	// Set up the HTTP server to listen on port 8080 and handle WebSocket connections
	http.HandleFunc("/", handleConnections)
	log.Info("Starting WebSocket server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
