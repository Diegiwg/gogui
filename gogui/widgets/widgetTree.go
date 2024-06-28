package gogui_widgets

import (
	"fmt"
	"strconv"
)

type WidgetTree struct {
	Widgets map[int]*Widget
}

func NewWidgetTree() *WidgetTree {
	return &WidgetTree{
		Widgets: make(map[int]*Widget),
	}
}

func (tree *WidgetTree) AddWidget(widget *Widget) int {
	tree.Widgets[len(tree.Widgets)+1] = widget
	return len(tree.Widgets)
}

func (tree *WidgetTree) GetWidget(id int) *Widget {
	return tree.Widgets[id]
}

func (tree *WidgetTree) Render() string {
	var html string = ""

	counter := len(tree.Widgets)
	if counter == 0 {
		return html
	}

	for i := 1; i <= counter; i++ {
		w := *tree.Widgets[i]
		html += fmt.Sprint(w.Html(strconv.Itoa(i)))
		html += "\n"
	}

	html = html[:len(html)-1]
	return html
}
