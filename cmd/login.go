/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/shailesz/cli-chat-golang/src/constants"
	"github.com/shailesz/cli-chat-golang/src/controllers"
	"github.com/shailesz/cli-chat-golang/src/helpers"
	"github.com/shailesz/cli-chat-golang/src/models"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := helpers.ReadConfig()

		controllers.Login(config)

		helpers.ClearScreen()
		helpers.WelcomeText()

		controllers.Socket.On("message", func(chat models.ChatMessage) {
			if chat.Username != config.Username {
				helpers.ClearLine()
				fmt.Println(constants.PURPLE_TERMINAL_COLOR + chat.ToString() + constants.RESET_TERMINAL_COLOR)
				helpers.Prompt()
			}
		})

		controllers.HandleChatInput(config)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

}
