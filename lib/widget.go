package lib

import (
	"fmt"
	"strings"

	"nhooyr.io/websocket"
)

type WidgetKind int

const (
	WidgetElement = iota
	WidgetLabel
	WidgetButton
)

func (wk *WidgetKind) String(w *Widget) string {
	switch *wk {
	case WidgetElement:
		return fmt.Sprintf("WidgetElement(%s)", w.GetData("tag").(string))
	case WidgetLabel:
		return "WidgetLabel"
	}
	return "WidgetKind"
}

type WidgetTree map[int]*Widget
type WidgetData map[string]interface{}
type WidgetRender func() string
type WidgetEvents map[string]EventHandler

type Widget struct {
	id       string
	kind     WidgetKind
	render   WidgetRender
	data     WidgetData
	children WidgetTree
	events   WidgetEvents
}

func newWidget() *Widget {
	return &Widget{
		data:     make(WidgetData),
		children: make(WidgetTree),
		events:   make(WidgetEvents),
	}
}

func (w *Widget) Dump(identLevel int) {
	println(fmt.Sprintf("%s%s", strings.Repeat(" ", identLevel), w.kind.String(w)))

	childCount := len(w.children)
	for i := 1; i <= childCount; i++ {
		w.children[i].Dump(identLevel + 1)
	}
}

func (w *Widget) SetData(key string, value interface{}) {
	w.data[key] = value
	w.emitContentUpdate()
}

func (w *Widget) GetData(key string) interface{} {
	return w.data[key]
}

func (w *Widget) AddChild(children ...*Widget) error {
	if len(children) == 0 {
		return fmt.Errorf("no children to add")
	}

	count := len(children)
	for i := 0; i < count; i++ {
		index := len(w.children) + 1
		w.children[index] = children[i]
	}

	dom.Register(w)
	return nil
}

func (w *Widget) GetChild(id int) *Widget {
	return nil
}

func (w *Widget) Html() string {
	childHtml := ""
	childCount := len(w.children)
	for i := 1; i <= childCount; i++ {
		childHtml += w.children[i].Html()
	}

	return fmt.Sprintf(w.render(), childHtml)
}

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

func (widget *Widget) emitContentUpdate() {
	if wsConn == nil || wsCtx == nil {
		return
	}

	wsConn.Write(*wsCtx, websocket.MessageText, []byte("update-element-content:"+widget.id+"|"+widget.Html()))
}
