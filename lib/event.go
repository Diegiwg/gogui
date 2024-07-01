package lib

type EventPayload struct {
	Id     string                 `json:"id"`
	Action string                 `json:"action"`
	Data   map[string]interface{} `json:"data"`
}

type EventHandler func(widget *Widget, payload *EventPayload)
