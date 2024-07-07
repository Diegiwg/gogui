package lib

func NewElement(tag string, content string) *Widget {
	w := newWidget()
	w.kind = WidgetElement

	w.SetData("tag", tag)
	w.SetData("content", content)

	return w
}
