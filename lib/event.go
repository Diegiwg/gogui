package lib

type EventPayload struct {
	Id     string                 `json:"id"`
	Action string                 `json:"action"`
	Data   map[string]interface{} `json:"data"`
}

type EventHandler func(widget *Widget, payload *EventPayload)

func (w *Widget) SetEvent(key string, handler EventHandler) {
	w.events[key] = handler
}

func (w *Widget) GetEvent(key string) *EventHandler {
	event, ok := w.events[key]

	if !ok {
		return nil
	}

	return &event
}

func (w *Widget) HasEvent(key string) bool {
	_, ok := w.events[key]
	return ok
}
