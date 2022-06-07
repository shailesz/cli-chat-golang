package services

import (
	"encoding/json"
	"io/ioutil"
)

// WriteConfig writes config to file.
func WriteConfig(data interface{}) {
	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("config.json", file, 0644)
}
