package lib

import "fmt"

type WidgetEvent struct {
	Name string `json:"name"`
}

func (w *Widget) Events() []WidgetEvent {
	return w.events
}

func (w *Widget) SetEvent(name string, handler EventHandler) {
	w.events = append(w.events, WidgetEvent{name})

	// tag-event-id
	key := fmt.Sprintf("%s-%s-%s", w.GetData("tag").(string), name, w.id)
	registerEvent(key, handler)
}
