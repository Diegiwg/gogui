package gogui_widgets

import "fmt"

func (w *Widget) Element(tag string, content string) *Widget {
	newWidget := NewWidget()

	newWidget.Kind = "Element"
	newWidget.SetData("tag", tag)
	newWidget.SetData("content", content)
	newWidget.Html = newWidget.elementHtml

	return newWidget
}

func (w *Widget) elementHtml(id string) string {
	tag := w.GetData("tag").(string)
	content := w.GetData("content").(string)

	return fmt.Sprintf("<%s id=\"%s\">%s</%s>", tag, id, content, tag)
}
