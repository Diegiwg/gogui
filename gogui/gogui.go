package gogui

import (
	"fmt"
	"net/http"

	widgets "github.com/Diegiwg/gogui/gogui/widgets"
)

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

type App struct {
	config     *AppConfig
	widgetTree *widgets.WidgetTree
	actions    map[string]func(ctx *HttpCtx)
}

func NewApp(config *AppConfig) (*App, error) {
	if config == nil {
		config = new(AppConfig)
	}

	err := config.check()
	if err != nil {
		return nil, err
	}

	return &App{
		config:     config,
		widgetTree: widgets.NewWidgetTree(),
		actions:    make(map[string]func(ctx *HttpCtx)),
	}, nil
}

func (a *App) Run() error {
	println("Server is running on http://" + a.config.serverAddress())

	http.HandleFunc("/", a.requestHandler)
	return http.ListenAndServe(a.config.serverAddress(), nil)
}

func (a *App) GetWidget(id int) *widgets.Widget {
	return a.widgetTree.GetWidget(id)
}
