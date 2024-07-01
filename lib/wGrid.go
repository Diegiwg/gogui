package lib

import "fmt"

func NewGrid(rows int, cols int) *Widget {
	w := newWidget()
	w.kind = WidgetLabel
	w.render = w.gridRender

	w.SetData("tag", "div")
	w.SetData("rows", rows)
	w.SetData("cols", cols)

	w.SetData("area", fmt.Sprintf(`class="grid-widget" style="grid-template-rows: repeat(%d, 1fr); grid-template-columns: repeat(%d, 1fr);"`, rows, cols))

	return w
}

func (w *Widget) gridRender() string {
	return w.renderOpenTag() + w.GetData("area").(string) + " >" + "%s" + w.renderCloseTag()
}
