package controllers

import (
	"log"

	"github.com/shailesz/cli-chat-golang-server/src/helpers"
	"github.com/shailesz/cli-chat-golang-server/src/services"
)

// SignUp creates a user.
func Signup(u, p string) {

	hp := helpers.Sha256(p)

	services.CreateUser(u, hp)

	log.Println("user created successfully!")
}
