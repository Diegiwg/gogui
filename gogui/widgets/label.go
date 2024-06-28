package gogui_widgets

import "fmt"

func (w *Widget) Label(text string) *Widget {
	w.Kind = "Label"
	w.SetData("text", text)
	w.Html = w.labelHtml

	return w
}

func (w *Widget) labelHtml(id string) string {
	content := w.GetData("text").(string)
	return fmt.Sprintf("<p id=\"%s\">%s</p>", id, content)
}
