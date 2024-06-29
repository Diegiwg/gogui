package gogui

import (
	"fmt"
)

// AppConfig is used to configure the application
type AppConfig struct {
	ServerAddr *string // default: 127.0.0.1
	ServerPort *int    // default: 8080
}

func NewConfig() *AppConfig {
	c := new(AppConfig)

	c.ServerAddr = new(string)
	c.ServerPort = new(int)

	return c
}

func (config *AppConfig) check() error {
	if config.ServerAddr == nil || *config.ServerAddr == "" {
		*config.ServerAddr = "127.0.0.1"
	}

	if config.ServerPort == nil || *config.ServerPort == 0 {
		*config.ServerPort = 8080
	}

	return nil
}

func (config *AppConfig) serverAddress() string {
	return fmt.Sprintf("%s:%d", *config.ServerAddr, *config.ServerPort)
}
