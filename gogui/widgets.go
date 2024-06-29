package gogui

import (
	"fmt"

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

// Button
//
// Button Params:
//
//	text string - text of the button
//	onClick func(ctx *HttpCtx, data HashMap) - function to call when the button is clicked
//
// Button Props:
//
//	disabled bool (default: false) - set to true to disable the button
//	secondary bool (default: false) - set to true to make a secondary button
func (app *App) Button(text string, onClick func(ctx *HttpCtx, data map[string]interface{})) (int, *W.Widget) {
	widget := W.NewWidget()
	widget = widget.Button(text)

	widgetId := app.widgetTree.AddWidget(widget)
	id := fmt.Sprintf("button-ID%d", widgetId)
	app.actions[id] = onClick

	return widgetId, widget
}

func (app *App) Grid(rows int, cols int) (int, *W.Widget) {
	widget := W.NewWidget()
	widget.Grid(rows, cols)

	return app.widgetTree.AddWidget(widget), widget
}
