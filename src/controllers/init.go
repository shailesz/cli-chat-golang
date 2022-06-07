package controllers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shailesz/cli-chat-golang/src/constants"
	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

// Conn is connection pool variable.
var Conn = InitDatabaseConnection()
var Socket = OpenConnection()

// InitDatabaseConnection initializes a database connection.
func InitDatabaseConnection() *pgxpool.Pool {

	// this returns connection pool
	conn, err := pgxpool.Connect(context.Background(), constants.DATABASE_URI)

	// error check
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

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
