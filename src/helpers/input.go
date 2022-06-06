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

func GetCredentials() (string, string) {
	fmt.Println("Please enter credentials.")

	u, p, err := Credentials()

	if err != nil {
		log.Panicln(err)
	}

	return u, p
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

func Prompt() {
	fmt.Print("-> You: ")
}

func ClearLine() {
	fmt.Fprint(os.Stdout, "\r \r")
}

func WelcomeText() {
	fmt.Println(constants.GREEN_TERMINAL_COLOR + "Welcome to chatroom! Type and press enter to chat..." + constants.RESET_TERMINAL_COLOR)
}
