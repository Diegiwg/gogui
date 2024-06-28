package gogui

import (
	gogui_widgets "github.com/Diegiwg/gogui/gogui/widgets"
)

func (app *App) Element(tag string, content string) (int, error) {
	widget := gogui_widgets.NewElement(tag, content)
	return app.widgetTree.AddWidget(widget), nil
}

func (app *App) Label(text string) (int, error) {
	widget := gogui_widgets.NewLabel(text)
	return app.widgetTree.AddWidget(widget), nil
}
