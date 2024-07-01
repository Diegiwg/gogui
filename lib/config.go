package lib

import (
	"fmt"
)

// Config is used to configure the application
type Config struct {
	ServerAddr *string // default: 127.0.0.1
	ServerPort *int    // default: 8080
}

func NewConfig() *Config {
	c := new(Config)

	c.ServerAddr = new(string)
	c.ServerPort = new(int)

	return c
}

func (config *Config) check() error {
	if config.ServerAddr == nil || *config.ServerAddr == "" {
		*config.ServerAddr = "127.0.0.1"
	}

	if config.ServerPort == nil || *config.ServerPort == 0 {
		*config.ServerPort = 8080
	}

	return nil
}

func (config *Config) serverAddress() string {
	return fmt.Sprintf("%s:%d", *config.ServerAddr, *config.ServerPort)
}
