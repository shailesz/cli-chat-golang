package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/shailesz/cli-chat-golang/cmd"
	"github.com/shailesz/cli-chat-golang/src/constants"
	"github.com/shailesz/cli-chat-golang/src/controllers"
	"github.com/shailesz/cli-chat-golang/src/helpers"
	"github.com/shailesz/cli-chat-golang/src/models"
)

var clear map[string]func() //create a map for storing clear funcs

func initConfig() {
	var config models.Config

	file, _ := json.MarshalIndent(config.Init(), "", " ")

	_ = ioutil.WriteFile("config.json", file, 0644)
}

func readConfig() models.Config {
	var jsonFile *os.File
	var err error
	var config models.Config

	for {
		// Open our jsonFile
		jsonFile, err = os.Open("config.json")
		// if os.Open returns an error then handle it
		if err != nil {
			initConfig()
		} else {
			byteValue, _ := ioutil.ReadAll(jsonFile)
			json.Unmarshal(byteValue, &config)

			jsonFile.Close()
			break
		}
	}

	return config
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func timeconv(timestamp int64) time.Time {
	t := time.Unix(0, timestamp)

	return t
}

func main() {

	// execute command
	cmd.Execute()

	config := readConfig()

	config.Login(controllers.Socket)

	CallClear()
	helpers.WelcomeText()

	controllers.Socket.On("message", func(chat models.ChatMessage) {
		if chat.Username != config.Username {
			helpers.ClearLine()
			fmt.Println(constants.PURPLE_TERMINAL_COLOR + chat.ToString() + constants.RESET_TERMINAL_COLOR)
			helpers.Prompt()
		}
	})

	controllers.HandleChatInput(config)

	defer controllers.Conn.Close()

}
