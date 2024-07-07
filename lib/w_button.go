package lib

func NewButton(text string, onClick EventHandler) *Widget {
	w := newWidget()
	w.kind = WidgetButton

	w.SetData("tag", "button")
	w.SetData("content", text)

	w.SetEvent("click", onClick)

	return w
}

// func (w *Widget) buttonRender() string {
// 	style := w.style.String()
// 	// TODO: implement dynamic event register instead of hardcoded onclick
// 	return w.renderOpenTag() + "onclick=\"buttonActionTrigger(this)\" style=\"" + style + "\" >" + w.GetData("text").(string) + "%s" + w.renderCloseTag()
// }
