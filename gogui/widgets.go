package gogui

import (
	"strconv"

	W "github.com/Diegiwg/gogui/gogui/widgets"
)

func (app *App) Element(tag string, content string) (int, *W.Widget) {
	widget := W.NewWidget()
	widget = widget.Element(tag, content)

	return app.widgetTree.AddWidget(widget), widget
}

func (app *App) Label(text string) (int, *W.Widget) {
	widget := W.NewWidget()
	widget = widget.Label(text)

	return app.widgetTree.AddWidget(widget), widget
}

func (app *App) Button(text string, onClick func(ctx *HttpCtx)) (int, *W.Widget) {
	widget := W.NewWidget()
	widget = widget.Button(text)

	widgetId := app.widgetTree.AddWidget(widget)
	app.actions["button-"+strconv.Itoa(widgetId)] = onClick

	return widgetId, widget
}

func (app *App) Grid(rows int, cols int) (int, *W.Widget) {
	widget := W.NewWidget()
	widget.Grid(rows, cols)

	return app.widgetTree.AddWidget(widget), widget
}
