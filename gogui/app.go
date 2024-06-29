package gogui

import (
	"net/http"

	widgets "github.com/Diegiwg/gogui/gogui/widgets"
)

type App struct {
	config     *AppConfig
	widgetTree *widgets.WidgetTree
	actions    map[string]func(ctx *HttpCtx, data map[string]interface{})
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
		actions:    make(map[string]func(ctx *HttpCtx, data map[string]interface{})),
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
