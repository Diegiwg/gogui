package lib

func NewButton(text string, onClick EventHandler) *Widget {
	w := newWidget()
	w.kind = WidgetButton

	w.SetData("tag", "button")
	w.SetData("content", text)

	w.SetEvent("click", nil, onClick)

	return w
}
