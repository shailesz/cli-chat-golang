package controllers

import (
	"bufio"
	"os"
	"time"

	"github.com/shailesz/cli-chat-golang/src/helpers"
	"github.com/shailesz/cli-chat-golang/src/models"
)

// HandleChatInput sends scanned input from to server.
func HandleChatInput(config models.Config) {
	reader := bufio.NewReader(os.Stdin)

	// prompt
	for {
		helpers.Prompt()
		data, _, _ := reader.ReadLine()
		message := string(data)
		SendChat(config.Username, message)
		if message == "$quit" {
			break
		}
	}
}

// SendChat emits chat event to server.
func SendChat(u, m string) {
	Socket.Emit("chat", models.ChatMessage{Username: u, Data: m, Timestamp: time.Now().UnixNano()})
}
