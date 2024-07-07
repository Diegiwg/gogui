package lib

import (
	"fmt"
	"strings"
)

type WidgetTree map[int]*Widget
type WidgetRender func(obj *RenderHtmlPayload)
type WidgetEvents map[string]EventHandler

type Widget struct {
	id       string
	kind     WidgetKind
	render   WidgetRender
	data     WidgetData
	children WidgetTree
	events   WidgetEvents
	style    WidgetStyle

	index  int
	parent *Widget
}

func newWidget() *Widget {
	return &Widget{
		data:     make(WidgetData),
		children: make(WidgetTree),
		events:   make(WidgetEvents),
		style:    NewWidgetStyle(),
	}
}

func (w *Widget) Delete() {
	// parent := w.parent
	dom.RemoveWidget(w.id)
	w.parent.RemoveChild(w.index)
	// parent.emitContentUpdate()
}

func (w *Widget) Dump(identLevel int) {
	println(fmt.Sprintf("%s%s", strings.Repeat(" ", identLevel), w.kind.String(w)))

	childCount := len(w.children)
	for i := 1; i <= childCount; i++ {
		w.children[i].Dump(identLevel + 1)
	}
}

func (w *Widget) AddChild(children ...*Widget) error {
	if len(children) == 0 {
		return fmt.Errorf("no children to add")
	}

	count := len(children)
	for i := 0; i < count; i++ {
		index := len(w.children) + 1
		w.children[index] = children[i]

		children[i].parent = w
		children[i].index = index
	}

	dom.Register(w)
	return nil
}

func (w *Widget) GetChild(index int) *Widget {
	if index < 1 || index > len(w.children) {
		return nil
	}

	return w.children[index]
}

func (w *Widget) RemoveChild(index int) {
	if index < 1 || index > len(w.children) {
		return
	}
	w.children[index] = nil
}

func (w *Widget) Render() *RenderHtmlPayload {
	var content string = ""
	if w.HasData("content") {
		content = w.GetData("content").(string)
	}

	obj := &RenderHtmlPayload{
		Id:         w.id,
		Tag:        w.GetData("tag").(string),
		Attributes: make([]WidgetAttribute, 0),
		Events:     w.Events(),
		Content:    content,
		Style:      w.style.String(),
	}

	if w.render != nil {
		w.render(obj)
	}

	obj.Children = make([]*RenderHtmlPayload, 0, len(w.children))
	for i := 0; i <= len(w.children); i++ {
		child := w.children[i]
		if child == nil {
			continue
		}
		obj.Children = append(obj.Children, child.Render())
	}

	return obj
}

// TODO: move to widgetAttribute.go
type WidgetAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// TODO: move to widgetEvent.go
type WidgetEvent struct {
	Name string `json:"name"`
}

func (w *Widget) Events() []WidgetEvent {
	events := make([]WidgetEvent, 0, len(w.events))
	for name, _ := range w.events {
		events = append(events, WidgetEvent{
			Name: name,
		})
	}
	return events
}

func (w *Widget) SetEvent(name string) {
	w.events[name] = nil
}
