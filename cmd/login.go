/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login logs in the user.",
	Long:  `Login logs in the user for given username and password.`,
	Run: func(cmd *cobra.Command, args []string) {
		u, p, _ := credentials()
		fmt.Println(u, p)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func credentials() (string, string, error) {
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
