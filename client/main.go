package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

// this is a simple client to talk to the gacman backend websocket

const (
	serverURL = "ws://localhost:8080" // The WebSocket server URL
)

func main() {
	// Connect to the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Errorf("Failed to close WebSocket connection: %v", err)
		}
	}()

	log.Infof("Connected to WebSocket server at %s", serverURL)

	// Start a goroutine to listen for messages from the server
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Errorf("Error reading message from server: %v", err)
				return
			}
			log.Infof("Received response from server: %s", string(message))
		}
	}()

	// Create a reader to read input from the console
	reader := bufio.NewReader(os.Stdin)

	// Main input loop
	for {
		fmt.Print("Enter direction (n, w, e, s): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Validate input
		if input != "n" && input != "w" && input != "e" && input != "s" {
			log.Warn("Invalid input. Please enter one of: n, w, e, s.")
			continue
		}

		// Send the direction to the server
		err := conn.WriteMessage(websocket.TextMessage, []byte(input))
		if err != nil {
			log.Errorf("Error sending message to server: %v", err)
			break
		}

		log.Infof("Sent direction: %s", input)
	}
}
