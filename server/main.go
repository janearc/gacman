package main

import (
	"gacman/core"
	"gacman/types"
	"net/http"

	"gacman/models"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

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

	// Initialize the space and get the starting coordinates
	space, startingCoord := models.InitSpace(10)

	// Retrieve the starting cell using the starting coordinates
	currentCell, exists := space.GetCell(startingCoord)
	if !exists {
		log.Errorf("Starting cell not found at coordinates: %s", startingCoord)
		return
	}

	// Send initial cell to the client
	initialData := currentCell.ToJSON()
	err = ws.WriteMessage(websocket.TextMessage, []byte(initialData))
	if err != nil {
		log.Errorf("Error sending initial data to client: %v", err)
		return
	}

	// Main loop: Listen for client input and respond with generated content
	for {
		// Wait for input from the client (e.g., "n", "w", "e", "s")
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Errorf("Error reading message from client: %v", err)
			break
		}

		direction := string(message)
		log.Infof("Received direction: %s", direction)

		// Generate a new cell in the specified direction
		newCell := types.GenerateNewCell(currentCell, direction)
		coord := core.GetCoordString(int(newCell.Position.X()), int(newCell.Position.Y()))

		// Add the new cell to the space if it's not already there
		if _, exists := space.GetCell(coord); !exists {
			space.AddCell(coord, newCell)
			currentCell = newCell // Update the current cell to the newly created one
		} else {
			// If the cell already exists, set the current cell to the existing one
			currentCell, _ = space.GetCell(coord)
		}

		// Send the new cell data back to the client
		response := currentCell.ToJSON()
		err = ws.WriteMessage(websocket.TextMessage, []byte(response))
		if err != nil {
			log.Errorf("Error sending response to client: %v", err)
			break
		}

		log.Infof("Sent cell data: %s", response)
	}
}

func main() {
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
