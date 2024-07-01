package lib

func NewElement(tag string, content string) *Widget {
	w := newWidget()
	w.kind = WidgetElement
	w.render = w.elementRender

	w.SetData("tag", tag)
	w.SetData("content", content)

	return w
}

func (w *Widget) elementRender() string {
	style := w.style.String()
	return w.renderOpenTag() + " style=\"" + style + "\" >" + w.GetData("content").(string) + "%s</" + w.GetData("tag").(string) + ">"
}
