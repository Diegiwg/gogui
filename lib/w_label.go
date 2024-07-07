package lib

func NewLabel(content string) *Widget {
	w := newWidget()
	w.kind = WidgetLabel

	w.SetData("tag", "p")
	w.SetData("content", content)

	return w
}
