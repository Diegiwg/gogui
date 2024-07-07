package lib

import (
	"encoding/json"
	"fmt"
	"log"

	"nhooyr.io/websocket"
)

type EventHandler func(widget *Widget, event *Event)

type Event struct {
	Id     string      `json:"id"`
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

var events = make(map[string]EventHandler)

func registerEvent(action string, handler EventHandler) {
	events[action] = handler
}

func handleEvent(event Event) {
	if event.Action == "" {
		return
	}

	handler := events[event.Action]
	if handler == nil {
		return
	}

	widget := widgetPool[event.Id]
	handler(widget, &event)
}

func emitEvent(eventKind string, data interface{}) {
	if wsConn == nil || wsCtx == nil {
		parsedData, _ := json.Marshal(data)

		log.Println(fmt.Sprintf("ERROR: Cannot emit event %s, no websocket connection. Data: %s", eventKind, parsedData))
		return
	}

	event := Event{
		Id:     "",
		Action: eventKind,
		Data:   data,
	}
	serializeEvent, _ := json.Marshal(event)

	wsConn.Write(*wsCtx, websocket.MessageText, []byte(serializeEvent))
}

type RenderHtmlPayload struct {
	Id         string               `json:"id"`
	Tag        string               `json:"tag"`
	Attributes []WidgetAttribute    `json:"attributes"`
	Events     []WidgetEvent        `json:"events"`
	Content    string               `json:"content"`
	Style      string               `json:"style"`
	Children   []*RenderHtmlPayload `json:"children"`
}

func emitRenderHtmlEvent(root *Widget) {
	emitEvent("render-html", root.Render())
}
