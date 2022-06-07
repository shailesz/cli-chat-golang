package controllers

import (
	"log"

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
		log.Panicln(err)
	}

	return client
}
