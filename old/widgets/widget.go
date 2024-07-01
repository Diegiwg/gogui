package gogui_widgets

import (
	"context"
	"fmt"

	"nhooyr.io/websocket"
)

type WsClient struct {
	WsConn *websocket.Conn
	Ctx    *context.Context
}

type HashMap map[string]interface{}

type Widget struct {
	Id    string
	Kind  string
	Html  func(id string) string
	Data  HashMap
	Style *WidgetStyle
	Props HashMap
}

func NewWidget() *Widget {
	return &Widget{
		Data:  make(HashMap),
		Style: &WidgetStyle{},
		Props: make(HashMap),
	}
}

func (widget *Widget) SetData(key string, value interface{}) {
	widget.Data[key] = value
}

func (widget *Widget) GetData(key string) interface{} {
	return widget.Data[key]
}

func (widget *Widget) SetProp(key string, value interface{}) {
	widget.Props[key] = value
}

func (widget *Widget) GetProp(key string) interface{} {
	return widget.Props[key]
}

func (widget *Widget) Child(newWidget *Widget) error {
	switch widget.Kind {
	case "Grid":
		children := widget.GetData("children").(map[int]*Widget)
		children[len(children)+1] = newWidget

		widget.SetData("children", children)

		return nil

	default:
		return fmt.Errorf("can't add child to widget of kind %s", widget.Kind)
	}
}

func (widget *Widget) EmitContentUpdate(wsConn *websocket.Conn, ctx *context.Context) {
	if wsConn == nil || ctx == nil {
		return
	}

	wsConn.Write(*ctx, websocket.MessageText, []byte("update-element-content:"+widget.Id+"|"+widget.Html(widget.Id)))
}
