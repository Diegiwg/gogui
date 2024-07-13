package lib

import "fmt"

type WidgetEvent struct {
	Name string   `json:"name"`
	Data []string `json:"data"`
}

func (w *Widget) Events() []WidgetEvent {
	return w.events
}

func (w *Widget) SetEvent(name string, targetData []string, handler EventHandler) {
	exists := w.HasEvent(name)
	if exists {
		w.RemoveEvent(name)
	}

	w.events = append(w.events, WidgetEvent{name, targetData})

	// tag-event-id
	key := fmt.Sprintf("%s-%s-%s", w.GetData("tag").(string), name, w.id)
	registerEvent(key, handler)
}

func (w *Widget) GetEvent(name string) EventHandler {
	key := fmt.Sprintf("%s-%s-%s", w.GetData("tag").(string), name, w.id)
	return events[key]
}

func (w *Widget) RemoveEvent(name string) {
	var newEvents []WidgetEvent
	for _, event := range w.events {
		if event.Name == name {
			continue
		}
		newEvents = append(newEvents, event)
	}
	w.events = newEvents
}

func (w *Widget) HasEvent(name string) bool {
	for _, event := range w.events {
		if event.Name == name {
			return true
		}
	}

	return false
}
