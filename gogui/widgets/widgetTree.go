package gogui_widgets

import (
	"fmt"
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

func (tree *WidgetTree) Render() (string, string) {
	var html string = ""
	var css string = ""

	counter := len(tree.Widgets)
	if counter == 0 {
		return html, css
	}

	for i := 1; i <= counter; i++ {
		w := *tree.Widgets[i]
		id := fmt.Sprintf("ID%d", i)
		html += w.Html(id) + "\n"
		css += w.Style.String(id) + "\n"
	}

	html = html[:len(html)-1]
	return html, css
}
