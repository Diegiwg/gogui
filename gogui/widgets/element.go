package gogui_widgets

import "fmt"

type Element struct {
	tag      string
	children map[int]*Widget
}

func (element *Element) Html(id string) string {
	var html string = fmt.Sprintf("<%s id=\"%s\">", element.tag, id)

	for childId, widget := range element.children {
		w := *widget
		html += fmt.Sprint(w.Html(
			fmt.Sprintf("e-%s-c-%d", id, childId),
		))
	}

	html += fmt.Sprintf("</%s>", element.tag)

	return html
}

func NewElement(tag string) *Element {
	return &Element{
		tag:      tag,
		children: make(map[int]*Widget),
	}
}

func (element *Element) addWidget(widget Widget) int {
	element.children[len(element.children)+1] = &widget
	return len(element.children)
}

// Supported Widgets //
func (element *Element) Label(text string) (int, error) {
	widget := NewLabel(text)
	return element.addWidget(widget), nil
}
