package lib

import (
	"fmt"
	"strings"
)

var widgetPool = make(map[string]*Widget, 0)

type WidgetTree map[int]*Widget
type WidgetRender func(obj *RenderHtmlPayload)

type Widget struct {
	id       string
	kind     WidgetKind
	render   WidgetRender
	data     WidgetData
	children WidgetTree
	events   []WidgetEvent
	style    WidgetStyle

	index  int
	parent *Widget
}

func (w *Widget) GetId() string {
	return w.id
}

func newWidget() *Widget {
	id := generateId(10)
	w := &Widget{
		id:       id,
		data:     make(WidgetData),
		children: make(WidgetTree),
		events:   make([]WidgetEvent, 0),
		style:    NewWidgetStyle(),
	}

	widgetPool[id] = w
	return w
}

func (w *Widget) Delete() {
	w.parent.RemoveChild(w.index)
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

	// dom.Register(w)
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

	child := w.children[index]

	delete(w.children, index)
	emitEvent("delete-widget", child.id)
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

// ATTRIBUTE SECTION //

type WidgetAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// EVENT SECTION //
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

// DATA SECTION //

type WidgetData map[string]interface{}

func (w *Widget) SetData(key string, value interface{}) {
	w.data[key] = value
}

func (w *Widget) GetData(key string) interface{} {
	return w.data[key]
}

func (w *Widget) HasData(key string) bool {
	_, ok := w.data[key]
	return ok
}

func (w *Widget) UpdateData(key string, value interface{}) {
	w.data[key] = value
	emitRenderHtmlEvent(w)
}

func (w *Widget) DeleteData(key string) {
	delete(w.data, key)
	// w.emitContentUpdate()
}

func (w *Widget) ClearData() {
	w.data = make(WidgetData)
	// w.emitContentUpdate()
}

// KIND SECTION //

type WidgetKind int

const (
	WidgetElement = iota
	WidgetLabel
	WidgetButton
	WidgetGrid
)

func (wk *WidgetKind) String(w *Widget) string {
	switch *wk {
	case WidgetElement:
		return fmt.Sprintf("WidgetElement(%s)", w.GetData("tag").(string))
	case WidgetLabel:
		return "WidgetLabel"
	case WidgetButton:
		return "WidgetButton"
	case WidgetGrid:
		return fmt.Sprintf("WidgetGrid(%d, %d)", w.GetData("rows").(int), w.GetData("cols").(int))
	}
	return "WidgetKind"
}

// STYLE SECTION //
type WidgetStyle struct {
	data map[string]interface{}
}

func NewWidgetStyle() WidgetStyle {
	return WidgetStyle{
		data: make(map[string]interface{}),
	}
}

func (w *WidgetStyle) Set(key string, value interface{}) {
	w.data[key] = value
}

func (w *WidgetStyle) Get(key string) interface{} {
	return w.data[key]
}

func (w *WidgetStyle) String() string {
	style := ""
	for key, value := range w.data {
		style += fmt.Sprintf("%s: %s; ", key, value)
	}

	return style
}

func (w *Widget) SetStyle(key string, value interface{}) {
	w.style.Set(key, value)
	// w.emitContentUpdate()
}

func (w *Widget) GetStyle(key string) interface{} {
	return w.style.Get(key)
}

func (w *Widget) HasStyle(key string) bool {
	_, ok := w.style.data[key]
	return ok
}

func (w *Widget) DeleteStyle(key string) {
	delete(w.style.data, key)
	// w.emitContentUpdate()
}
