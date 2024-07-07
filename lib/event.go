package lib

import (
	"context"
	"encoding/json"

	"nhooyr.io/websocket"
)

type EventHandler func(widget *Widget, event *Event, conn *websocket.Conn, ctx *context.Context)

type Event struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

var events = make(map[string]EventHandler)

func registerEvent(action string, handler EventHandler) {
	events[action] = handler
}

func handleEvent(event Event, conn *websocket.Conn, ctx *context.Context) {
	if event.Action == "" {
		return
	}

	handler := events[event.Action]
	if handler == nil {
		return
	}

	handler(nil, &event, conn, ctx)
}

func emitEvent(eventKind string, data interface{}, conn *websocket.Conn, ctx *context.Context) {
	event := Event{eventKind, data}
	serializeEvent, _ := json.Marshal(event)

	conn.Write(*ctx, websocket.MessageText, []byte(serializeEvent))
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

func emitRenderHtmlEvent(root *Widget, conn *websocket.Conn, ctx *context.Context) {
	emitEvent("render-html", root.Render(), conn, ctx)
}
