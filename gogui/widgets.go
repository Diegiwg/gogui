package gogui

import (
	"strconv"

	gogui_widgets "github.com/Diegiwg/gogui/gogui/widgets"
)

func (app *App) Element(tag string, content string) (int, error) {
	widget := gogui_widgets.NewElement(tag, content)
	return app.widgetTree.AddWidget(widget), nil
}

func (app *App) Label(text string) (*gogui_widgets.Label, int, error) {
	widget := gogui_widgets.NewLabel(text)
	return widget, app.widgetTree.AddWidget(widget), nil
}

func (app *App) Button(text string, onClick func(ctx *HttpCtx)) (int, error) {
	widget := gogui_widgets.NewButton(text)
	id := app.widgetTree.AddWidget(widget)
	app.actions["button-"+strconv.Itoa(id)] = onClick
	return id, nil
}
