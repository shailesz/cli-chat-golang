package controllers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/shailesz/cli-chat-golang/src/helpers"
)

// CreateUser creates a user.
func CreateUser(u, p string) {

	hp := helpers.Sha256(p)

	rows := [][]interface{}{
		{u, hp},
	}

	copyCount, err := Conn.CopyFrom(
		context.TODO(),
		pgx.Identifier{"users"},
		[]string{"username", "password"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create user: %v\n", err)
		os.Exit(1)
	}

	log.Println("user created successfully!", copyCount)
}
