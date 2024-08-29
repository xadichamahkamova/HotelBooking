package handler

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	consumer "websocket-service/internal/kafka/consumer"
	"websocket-service/logger"

	"github.com/gorilla/websocket"
	"github.com/twmb/franz-go/pkg/kgo"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	clients    = make(map[*websocket.Conn]bool)
	clientsMux sync.Mutex
)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("Failed to upgrade WebSocket connection: ", err)
		return
	}
	defer func() {
		conn.Close()
		logger.Info("WebSocket connection closed")
	}()

	clientsMux.Lock()
	clients[conn] = true
	clientsMux.Unlock()
	logger.Info("New WebSocket client connected")

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			logger.Error("Failed to read message from WebSocket: ", err)
			clientsMux.Lock()
			delete(clients, conn)
			clientsMux.Unlock()
			logger.Info("WebSocket client disconnected")
			break
		}
	}
}

func ConsumeMessage(consumer consumer.ConsumeInit) error {

	ctx := context.Background()

	for {
		fetches := consumer.Client.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			logger.Error("Error consuming messages: ", errs)
			return fmt.Errorf("error consuming messages: %v", errs)
		}

		fetches.EachPartition(func(partition kgo.FetchTopicPartition) {
			for _, record := range partition.Records {
				logger.Info(fmt.Sprintf("Received message - Key: %s, Value: %s", string(record.Key), string(record.Value)))
				BroadcastToClients(string(record.Value))
			}
		})
	}
}

func BroadcastToClients(message string) {

	clientsMux.Lock()
	defer clientsMux.Unlock()

	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			logger.Error("Failed to send message to WebSocket client: ", err)
			client.Close()
			delete(clients, client)
		} else {
			logger.Info("Message sent to WebSocket client")
		}
	}
}
