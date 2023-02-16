package sse

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"social-network/internal/entities"

	"github.com/golang-jwt/jwt"
)

type MessageNotification struct {
	AccountID string `json:"account_id"`
	Message   string `json:"message"`
}

func (m *MessageNotification) Marshal() ([]byte, error) {
	msgByte, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	return msgByte, nil
}

func (m *MessageNotification) Unmarshal(val []byte) error {
	if err := json.Unmarshal(val, m); err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	return nil
}

type Broker struct {

	// Events are pushed to this channel by the main events-gathering routine
	Notifier chan []byte

	// New client connections
	newClients chan chan []byte

	// Closed client connections
	closingClients chan chan []byte

	// Client connections registry
	clients map[chan []byte]bool

	users map[string]chan []byte

	jwtKey string
}

func NewServer(jwtKey string) (broker *Broker) {
	// Instantiate a broker
	broker = &Broker{
		Notifier:       make(chan []byte, 1),
		newClients:     make(chan chan []byte),
		closingClients: make(chan chan []byte),
		clients:        make(map[chan []byte]bool),
		users:          make(map[string]chan []byte),
		jwtKey:         jwtKey,
	}

	// Set it running - listening and broadcasting events
	go broker.listen()

	return
}

func (broker *Broker) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	flusher, ok := rw.(http.Flusher)
	if !ok {
		http.Error(rw, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/event-stream")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	// Each connection registers its own message channel with the Broker's connections registry
	messageChan := make(chan []byte)

	token := req.URL.Query().Get("token")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(broker.jwtKey), nil
	})
	if err != nil {
		http.Error(rw, fmt.Sprintf("token isn't valid 1 %v %v", err, token), http.StatusBadRequest)
		return
	}
	_ = t

	accountID, ok := claims["id"].(string)
	if !ok {
		http.Error(rw, "token isn't valid 2", http.StatusBadRequest)
		return
	}

	// Signal the broker that we have a new connection
	broker.users[accountID] = messageChan
	broker.newClients <- messageChan

	// Remove this client from the map of connected clients
	// when this handler exits.
	defer func() {
		broker.closingClients <- messageChan
	}()

	// Listen to connection close and un-register messageChan
	// notify := rw.(http.CloseNotifier).CloseNotify()
	notify := req.Context().Done()

	go func() {
		<-notify
		broker.closingClients <- messageChan
	}()

	for {
		// Write to the ResponseWriter
		// Server Sent Events compatible
		fmt.Fprintf(rw, "data: %s\n\n", <-messageChan)

		// Flush the data immediately instead of buffering it for later.
		flusher.Flush()
	}
}

func (broker *Broker) listen() {
	for {
		select {
		case s := <-broker.newClients:

			// A new client has connected.
			// Register their message channel
			broker.clients[s] = true
			log.Printf("Client added. %d registered clients", len(broker.clients))
		case s := <-broker.closingClients:

			// A client has detached and we want to
			// stop sending them messages.
			delete(broker.clients, s)
			log.Printf("Removed client. %d registered clients", len(broker.clients))
		case message := <-broker.Notifier:
			messageNotification := &MessageNotification{}
			messageNotification.Unmarshal(message)

			clientMessageChan, ok := broker.users[messageNotification.AccountID]
			if ok {
				clientMessageChan <- []byte(messageNotification.Message)
			}
		}
	}
}

func UnmarshalFeed(val []byte) (*entities.Feed, error) {
	var msg entities.Feed
	if err := json.Unmarshal(val, &msg); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return &msg, nil
}

/*
user1
var client = new EventSource("http://localhost:8082?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzY1MjU0MTEsImlkIjoiY2Ztb3Rhc2R0Nm1yc3BvZ3U2NTAiLCJ1c2VybmFtZSI6InVzZXJuYW1lMSJ9.w2tyI7-FAiybO0r2EXbf90sHptcV3ZnfLtltLbuNx6c")
client.onmessage = function (msg) {
  console.log(msg)
}

user2
var client = new EventSource("http://localhost:8082?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzY1MjM1NDEsImlkIjoiY2Ztb3Voa2R0Nm1yc3BvZ3U2NWciLCJ1c2VybmFtZSI6InVzZXJuYW1lMiJ9.RlwydZ23-fNV25hTDf3rgBXkOQwK5u8FStNNDE7pxQg")
client.onmessage = function (msg) {
  console.log(msg)
}

user3
var client = new EventSource("http://localhost:8082?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzY1MjM2MTEsImlkIjoiY2ZtcGZ2NGR0Nm1yc3BvZ3U2NmciLCJ1c2VybmFtZSI6InVzZXJuYW1lMyJ9.48hcOvUimVcEE-EAOYpzs2cfadAOwvb1xQ99_N9jQAw")
client.onmessage = function (msg) {
  console.log(msg)
}

*/
