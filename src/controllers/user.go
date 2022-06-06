package controllers

import (
	"fmt"

	"github.com/shailesz/cli-chat-golang/src/models"
	"github.com/shailesz/cli-chat-golang/src/services"
)

// CreateUser creates a user.
func CreateUser(u, p string) {
	var waitResponse bool = true

	user := models.User{Username: u, Password: p}

	Socket.On("signup", func(res models.AuthMessage) {
		if res.Status == 200 {
			fmt.Println("Successfully signed up, please continue to login.")
			services.WriteConfig(user)
		} else {
			fmt.Println("Something went wrong! Please try again.")
		}
		waitResponse = false
	})

	Socket.Emit("signup", user)

	for {
		if !waitResponse {
			break
		}
	}
}
