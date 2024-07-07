package lib

import "fmt"

func NewGrid(rows int, cols int) *Widget {
	w := newWidget()
	w.kind = WidgetLabel
	w.render = w.gridRender

	w.SetData("tag", "div")
	w.SetData("rows", rows)
	w.SetData("cols", cols)

	w.SetStyle("grid-template-rows", fmt.Sprintf("repeat(%d, 1fr)", rows))
	w.SetStyle("grid-template-columns", fmt.Sprintf("repeat(%d, 1fr)", cols))

	return w
}

func (w *Widget) gridRender(obj *RenderHtmlPayload) {
	obj.Attributes = append(obj.Attributes, WidgetAttribute{"class", "grid-widget"})
}
