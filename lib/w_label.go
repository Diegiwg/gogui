package lib

func NewLabel(content string) *Widget {
	w := newWidget()
	w.kind = WidgetLabel
	w.render = w.labelRender

	w.SetData("tag", "p")
	w.SetData("content", content)

	return w
}

func (w *Widget) labelRender() string {
	return w.renderOpenTag() + ">" + w.GetData("content").(string) + "%s" + w.renderCloseTag()
}
