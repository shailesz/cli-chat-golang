package main

import (
	"github.com/shailesz/cli-chat-golang/cmd"
	"github.com/shailesz/cli-chat-golang/src/controllers"
)

func main() {

	// execute command
	cmd.Execute()

	// defer database connection close
	defer controllers.Conn.Close()
}
