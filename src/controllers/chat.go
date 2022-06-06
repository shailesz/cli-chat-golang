package controllers

import (
	"bufio"
	"os"
	"time"

	"github.com/shailesz/cli-chat-golang/src/helpers"
	"github.com/shailesz/cli-chat-golang/src/models"
)

func HandleChatInput(config models.Config) {
	reader := bufio.NewReader(os.Stdin)

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

func SendChat(u, m string) {
	Socket.Emit("chat", models.ChatMessage{Username: u, Data: m, Timestamp: time.Now().UnixNano()})
}
