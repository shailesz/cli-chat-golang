package services

import (
	"context"
	"fmt"
	"os"

	"github.com/JackC/pgx"
)

func CreateUser(u, p string) {

	rows := [][]interface{}{
		{u, p},
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

}
