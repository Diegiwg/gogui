package gogui_widgets

import "fmt"

func (w *Widget) Button(text string) *Widget {
	w.Kind = "Button"
	w.SetData("text", text)
	w.Html = w.buttonHtml

	return w
}

func (w *Widget) buttonHtml(id string) string {
	text := w.GetData("text").(string)
	return fmt.Sprintf("<button id=\"%s\" onclick=\"buttonActionTrigger(this)\">%s</button>", id, text)
}
