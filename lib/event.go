package lib

import (
	"context"
	"encoding/json"
	"log"

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
	log.Println("handleEvent")

	handler := events[event.Action]
	if handler == nil {
		log.Println("event not found: " + event.Action)
		return
	}

	handler(nil, &event, conn, ctx)
}

func emitEvent(eventKind string, data interface{}, conn *websocket.Conn, ctx *context.Context) {
	serializeData, _ := json.Marshal(data)
	log.Println("emitEvent: " + string(serializeData))

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
