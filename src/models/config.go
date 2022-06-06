package models

import (
	"github.com/shailesz/cli-chat-golang/src/services"
)

// Config type for config file.
type Config User

// Init initializes config.
func (c *Config) Init() Config {
	return Config{Username: "", Password: ""}
}

// Update updates config file with given parameters.
func (c *Config) Update(u, p string) Config {

	c.Username, c.Password = u, p

	// write config file.
	services.WriteConfig(c)

	return Config{Username: u, Password: p}
}
