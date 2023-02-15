package sse

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"social-network/internal/entities"

	"github.com/golang-jwt/jwt"
)

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

			// A client has dettached and we want to
			// stop sending them messages.
			delete(broker.clients, s)
			log.Printf("Removed client. %d registered clients", len(broker.clients))
		case event := <-broker.Notifier:

			// We got a new event from the outside!
			// Send event to all connected clients
			// for clientMessageChan, _ := range broker.clients {
			// 	clientMessageChan <- event
			// }

			feed, _ := UnmarshalFeed(event)
			clientMessageChan, ok := broker.users[feed.AccountID]
			if ok {
				clientMessageChan <- []byte(fmt.Sprintf("account(%s) create new feed with message(%s)", feed.AccountID, feed.Message))
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
var client = new EventSource("http://localhost:8082?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzY0OTEwNTYsImlkIjoiY2ZtNTdpc2R0Nm1paTB1cmMzN2ciLCJ1c2VybmFtZSI6InVzZXJuYW1lMSJ9.w5EN6zUGdJTN9tF_XoMsFua-oSnqyzh-Exs2gzPLOcs")
client.onmessage = function (msg) {
  console.log(msg)
}

*/
