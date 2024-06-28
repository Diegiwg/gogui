package gogui_widgets

import "fmt"

type Label struct {
	text string
}

func (label *Label) Html(id string) string {
	return fmt.Sprintf("<p id=\"%s\">%s</p>", id, label.text)
}

func NewLabel(text string) *Label {
	return &Label{
		text: text,
	}
}

func (label *Label) SetText(text string) {
	label.text = text
}

func (label *Label) GetText() string {
	return label.text
}
