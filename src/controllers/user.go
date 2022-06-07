package controllers

import (
	"fmt"

	"github.com/shailesz/cli-chat-golang/src/helpers"
	"github.com/shailesz/cli-chat-golang/src/models"
	"github.com/shailesz/cli-chat-golang/src/services"
)

// CreateUser creates a user.
func CreateUser(e, u, p string) {
	var waitResponse bool = true

	user := models.User{Email: e, Username: u, Password: p}

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

// Login logs in user from config file.
func Login(c models.Config) (string, string) {
	var isWaiting, isUpdate bool
	var u, p string

	// handle configs from config file
	if c.Username == "" || c.Password == "" {
		_, u, p = helpers.GetCredentials(false)

		isUpdate = true

	} else {
		fmt.Println("Processing...")
		u, p = c.Username, c.Password
		isUpdate = false
	}

	// listener for auth messages.
	Socket.On("auth", func(message models.AuthMessage) {
		if message.Status == 404 {
			fmt.Println("You could not be authenticated. please try again.")
		} else {
			fmt.Println("Authenticated.")

			if isUpdate {
				c.Update(u, p)
			}
		}

		isWaiting = false
	})

	isWaiting = true
	Socket.Emit("auth", models.User{Username: u, Password: p}) // send auth message to server

	// wait for auth message.
	for {
		if !isWaiting {
			break
		}
	}

	return u, p
}
