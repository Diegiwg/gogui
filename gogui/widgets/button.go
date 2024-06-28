package gogui_widgets

import "fmt"

type Button struct {
	text string
}

func (button *Button) Html(id string) string {
	return fmt.Sprintf("<button id=\"%s\" onclick=\"buttonActionTrigger(this)\">%s</button>", id, button.text)
}

func NewButton(text string) *Button {
	return &Button{
		text: text,
	}
}

func (button *Button) SetText(text string) {
	button.text = text
}

func (button *Button) GetText() string {
	return button.text
}
