package gogui_widgets

import (
	"fmt"
)

// Button
//
// Button Data:
//
//	text string - text of the button
//
// Button Props:
//
//	disabled bool (default: false) - set to true to disable the button
//	secondary bool (default: false) - set to true to make a secondary button
func (w *Widget) Button(text string) *Widget {
	w.Kind = "Button"
	w.initButtonProps()
	w.SetData("text", text)
	w.Html = w.buttonHtml
	w.Style.Data = make(WidgetStyleData)

	return w
}

func (w *Widget) buttonHtml(id string) string {
	text := w.GetData("text").(string)

	var props string = ""
	if w.GetProp("disabled").(bool) {
		props = "disabled"
	}

	if w.GetProp("secondary").(bool) {
		setSecondaryStyle(&w.Style.Data)
	} else {
		setPrimaryStyle(&w.Style.Data)
	}

	return fmt.Sprintf("<button id=\"%s\" %s onclick=\"buttonActionTrigger(this)\">%s</button>", id, props, text)
}

func setPrimaryStyle(object *WidgetStyleData) {
	(*object)["background-color"] = "#04AA6D"
	(*object)["border"] = "none"
	(*object)["color"] = "white"
	(*object)["padding"] = "15px 32px"
	(*object)["text-align"] = "center"
	(*object)["text-decoration"] = "none"
	(*object)["display"] = "inline-block"
	(*object)["font-size"] = "16px"
}

func setSecondaryStyle(object *WidgetStyleData) {
	(*object)["background-color"] = "#f44336"
	(*object)["border"] = "none"
	(*object)["color"] = "white"
	(*object)["padding"] = "15px 32px"
	(*object)["text-align"] = "center"
	(*object)["text-decoration"] = "none"
	(*object)["display"] = "inline-block"
	(*object)["font-size"] = "16px"
}

func (w *Widget) initButtonProps() {
	w.SetProp("disabled", false)
	w.SetProp("secondary", false)
}
