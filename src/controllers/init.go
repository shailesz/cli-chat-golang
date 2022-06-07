package controllers

import (
	"log"
	"os"

	"github.com/shailesz/cli-chat-golang/src/constants"
	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

var Socket = OpenConnection()

// OpenConnection opens a websocket connection to server.
func OpenConnection() *socketio_client.Client {

	// options for socket
	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}

	// new client socket
	client, err := socketio_client.NewClient(constants.WEBSOCKET_URI, opts)
	if err != nil {
		log.Println("Could not connect to server! Please try again.")
		os.Exit(1)
	}

	return client
}
