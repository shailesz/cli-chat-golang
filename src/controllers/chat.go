package controllers

import (
	"log"

	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

// OpenConnection opens a websocket connection to server.
func OpenConnection() *socketio_client.Client {

	uri := "http://localhost:8000/socket.io/"

	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		log.Panicln(err)
	}

	return client

}
