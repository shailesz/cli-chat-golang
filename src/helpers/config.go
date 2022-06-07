package helpers

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/shailesz/cli-chat-golang/src/models"
)

var clear map[string]func() //create a map for storing clear funcs

// initConfig initializes an empty config file.
func initConfig() {
	var config models.Config

	file, _ := json.MarshalIndent(config.Init(), "", " ")

	_ = ioutil.WriteFile("config.json", file, 0644)
}

// ReadConfig reads from a config file.
func ReadConfig() models.Config {
	var jsonFile *os.File
	var err error
	var config models.Config

	for {
		// Open our config file
		jsonFile, err = os.Open("config.json")
		// if os.Open returns an error then handle it
		if err != nil {
			initConfig()
		} else {
			byteValue, _ := ioutil.ReadAll(jsonFile)
			json.Unmarshal(byteValue, &config)

			jsonFile.Close() // file already parsed, closing file
			break
		}
	}

	return config
}

// ClearScreen clears the terminal screen.
func ClearScreen() {
	clear = make(map[string]func()) //Initialize it

	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {
		value() // execute clear screen
	} else { // unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func Timeconv(timestamp int64) time.Time {
	t := time.Unix(0, timestamp)

	return t
}
