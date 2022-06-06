/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/shailesz/cli-chat-golang/src/controllers"
	"github.com/shailesz/cli-chat-golang/src/helpers"
	"github.com/spf13/cobra"
)

// signupCmd represents the signup command
var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "Create a new user",
	Long:  `Create a new user and get logged in to chat`,
	Run: func(cmd *cobra.Command, args []string) {
		u, p := helpers.GetCredentials()
		controllers.CreateUser(u, p)

	},
}

func init() {
	rootCmd.AddCommand(signupCmd)

}
