package gogui_widgets

import "fmt"

type Element struct {
	tag     string
	content string
}

func (element *Element) Html(id string) string {
	return fmt.Sprintf("<%s id=\"%s\">%s</%s>", element.tag, id, element.content, element.tag)
}

func NewElement(tag string, content string) *Element {
	return &Element{
		tag:     tag,
		content: content,
	}
}
