package lib

import (
	"fmt"
	"strings"
)

type WidgetTree map[int]*Widget
type WidgetRender func() string
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
	parent := w.parent
	dom.RemoveWidget(w.id)
	w.parent.RemoveChild(w.index)
	parent.emitContentUpdate()
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

func (w *Widget) Html() string {
	childHtml := ""
	childCount := len(w.children)
	for i := 1; i <= childCount; i++ {
		if w.children[i] == nil {
			continue
		}
		childHtml += w.children[i].Html()
	}

	return fmt.Sprintf(w.render(), childHtml)
}
