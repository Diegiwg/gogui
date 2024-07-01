package lib

import "log"

type DOM struct {
	widgets map[string]*Widget
}

func NewDOM() *DOM {
	return &DOM{
		widgets: make(map[string]*Widget),
	}
}

var dom = NewDOM()

func (dom *DOM) AddWidget(widget *Widget) error {
	var id string
	for {
		id = widget.generateId(10)

		if _, ok := dom.widgets[id]; !ok {
			break
		}
	}

	widget.id = id
	dom.widgets[widget.id] = widget
	return nil
}

func (dom *DOM) GetWidget(id string) *Widget {
	return dom.widgets[id]
}

func (dom *DOM) RemoveWidget(id string) error {
	delete(dom.widgets, id)
	return nil
}

func (dom *DOM) Register(root *Widget) {
	if root == nil {
		log.Fatal("root widget is nil")
	}

	if dom.widgets[root.id] == nil {
		dom.AddWidget(root)
	}

	for _, widget := range root.children {
		dom.Register(widget)
	}
}
