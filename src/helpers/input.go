package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/shailesz/cli-chat-golang/src/constants"
	"golang.org/x/term"
)

// GetCredentials gets credentials from user for signup or login.
func GetCredentials(isSignup bool) (string, string, string) {
	var e string = ""

	fmt.Println("Please enter credentials.")

	if isSignup {
		e, _ = GetEmail()
	}
	u, p, err := Credentials()

	if err != nil {
		log.Panicln(err)
	}

	return e, u, p
}

// GetEmail gets email from user.
func GetEmail() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(email), nil
}

// Credentials handles input of user credentials during login or signup.
func Credentials() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", err
	}

	password := string(bytePassword)
	fmt.Println()
	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}

// Prompt prints a prompt to terminal.
func Prompt() {
	fmt.Print("-> You: ")
}

// ClearLine clears the current line.
func ClearLine() {
	fmt.Fprint(os.Stdout, "\r \r")
}

// WelcomeText prints welcome text to terminal.
func WelcomeText() {
	fmt.Println(constants.GREEN_TERMINAL_COLOR + "Welcome to chatroom! Type and press enter to chat..." + constants.RESET_TERMINAL_COLOR)
}
