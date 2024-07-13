package lib

func NewInput() *Widget {
	w := newWidget()
	w.kind = WidgetInput

	w.SetData("tag", "button")

	return w
}
