package lib

// TODO: implement a way for user bind handlers for specific events, e.g. keyup:enter

func NewInput(bindValue *string) *Widget {
	w := newWidget()
	w.kind = WidgetInput

	w.SetData("tag", "input")
	w.SetData("value", bindValue)

	w.SetStyle("width", "100%")

	w.SetAttribute("type", "text")
	w.SetAttribute("value", *bindValue)

	w.SetEvent("keyup", []string{
		"key",
		"target.value",
	}, func(widget *Widget, event *Event) {
		data, ok := event.Data.(map[string]interface{})
		if !ok {
			return
		}

		// key := data["key"].(string)
		value := data["target.value"].(string)
		*bindValue = value
	})

	return w
}
