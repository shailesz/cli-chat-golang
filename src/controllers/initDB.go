package controllers

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Conn is connection pool variable.
var Conn = InitDatabaseConnection()

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
