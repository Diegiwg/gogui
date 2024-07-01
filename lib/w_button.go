// TODO: refactor this code
package lib

func NewButton(text string, onClick EventHandler) *Widget {
	w := newWidget()
	w.kind = WidgetButton
	w.render = w.buttonRender

	w.SetData("tag", "button")
	w.SetData("text", text)

	w.SetEvent("click", onClick)

	return w
}

func (w *Widget) buttonRender() string {
	return w.renderOpenTag() + "onclick=\"buttonActionTrigger(this)\" >" + w.GetData("text").(string) + "%s" + w.renderCloseTag()
}
