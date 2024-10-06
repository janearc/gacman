package main

import (
	"encoding/json"
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

		// Generate the new cell and check for movement status
		newCell, status := models.GenerateNewCell(currentCell, direction, &space)

		// If movement is successful, update the current cell
		if status == "Movement successful" {
			currentCell = newCell
		}

		// Prepare the response with the new cell data and the status message
		response := map[string]string{
			"cell":   currentCell.ToJSON(),
			"status": status,
		}

		// Serialize the response to JSON
		responseData, err := json.Marshal(response)
		if err != nil {
			log.Errorf("Error serializing response to JSON: %v", err)
			break
		}

		// Send the response back to the client
		err = ws.WriteMessage(websocket.TextMessage, responseData)
		if err != nil {
			log.Errorf("Error sending response to client: %v", err)
			break
		}

		log.Infof("Sent response: %s", responseData)
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
