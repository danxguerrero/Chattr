package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader  = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan string)
	mu        sync.Mutex
)

// Add this struct to parse the incoming JSON
type WSMessage struct {
	Message string                 `json:"message"`
	Headers map[string]interface{} `json:"HEADERS"`
}

func HandleWebSocket(c echo.Context) error {
	fmt.Println("WebSocket connection attempt...")
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Printf("Error upgrading connection: %v\n", err)
		return err
	}
	defer conn.Close()

	fmt.Println("WebSocket connection established successfully")
	mu.Lock()
	clients[conn] = true
	fmt.Printf("Client connected. Total clients: %d\n", len(clients))
	mu.Unlock()

	for {
		fmt.Println("Waiting for message...")
		_, msg, err := conn.ReadMessage()
		if err != nil {
			mu.Lock()
			delete(clients, conn)
			mu.Unlock()
			fmt.Printf("Error reading message: %v\n", err)
			break
		}

		// Parse the JSON message
		var wsMsg WSMessage
		if err := json.Unmarshal(msg, &wsMsg); err != nil {
			fmt.Printf("Error parsing message: %v\n", err)
			continue
		}

		fmt.Printf("Received message: %s\n", wsMsg.Message)
		messageHTML := fmt.Sprintf(`
			<div class="chat-message" hx-swap-oob="afterbegin:#chat-list">
				<p>%s</p>
			</div>`, template.HTMLEscapeString(wsMsg.Message))
		fmt.Printf("Sending message: %s\n", messageHTML)
		broadcast <- messageHTML
	}
	return nil
}

func BroadcastMessages() {
	for {
		msg := <-broadcast
		mu.Lock()
		fmt.Printf("Broadcasting to %d clients\n", len(clients))
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				fmt.Printf("Error writing to client: %v\n", err)
				client.Close()
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}
