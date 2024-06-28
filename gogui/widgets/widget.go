package gogui_widgets

import (
	"fmt"
	"strconv"
)

type Widget interface {
	Html(id string) string
}

type WidgetTree struct {
	Widgets map[int]*Widget
}

func NewWidgetTree() *WidgetTree {
	return &WidgetTree{
		Widgets: make(map[int]*Widget),
	}
}

func (tree *WidgetTree) AddWidget(widget Widget) int {
	tree.Widgets[len(tree.Widgets)+1] = &widget
	return len(tree.Widgets)
}

func (tree *WidgetTree) GetWidget(id int) *Widget {
	return tree.Widgets[id]
}

func (tree *WidgetTree) Render() string {
	var html string = ""

	for id, widget := range tree.Widgets {
		w := *widget
		html += fmt.Sprint(w.Html(strconv.Itoa(id)))
	}

	return html
}
