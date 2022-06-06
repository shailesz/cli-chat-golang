package controllers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

// Conn is connection pool variable.
var Conn = InitDatabaseConnection()
var Socket = OpenConnection()

// InitDatabaseConnection initializes a database connection.
func InitDatabaseConnection() *pgxpool.Pool {
	// hardcoded db url
	databaseUrl := "postgres://shailesz:password@localhost:5432/cli-chat-golang"

	// this returns connection pool
	conn, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

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
