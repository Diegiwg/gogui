package gogui

import (
	gogui_widgets "github.com/Diegiwg/gogui/gogui/widgets"
)

func (app *App) Element(tag string) (*gogui_widgets.Element, int, error) {
	widget := gogui_widgets.NewElement(tag)
	id := app.widgetTree.AddWidget(widget)
	return widget, id, nil
}

func (app *App) Label(text string) (int, error) {
	widget := gogui_widgets.NewLabel(text)
	return app.widgetTree.AddWidget(widget), nil
}
