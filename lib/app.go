package lib

import (
	"log"
	"net/http"
)

type App struct {
	Root *Widget

	config *Config
}

func NewApp(config *Config) (*App, error) {
	if config == nil {
		config = NewConfig()
	}

	err := config.check()
	if err != nil {
		return nil, err
	}

	return &App{
		Root: NewElement("div", ""),

		config: config,
	}, err
}

func (a *App) Dump() {
	a.Root.Dump(0)
}

func (a *App) Run() error {
	dom.Register(a.Root)

	log.Println("INFO: Server is running on http://" + a.config.serverAddress())

	http.HandleFunc("/", a.requestHandler)
	return http.ListenAndServe(a.config.serverAddress(), nil)
}
