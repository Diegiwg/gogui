package lib

import "fmt"

func (w *Widget) AddChild(children ...*Widget) error {
	if len(children) == 0 {
		return fmt.Errorf("no children to add")
	}

	count := len(children)
	for i := 0; i < count; i++ {
		child := children[i]

		w.children[child.id] = child
		w.childrenIds[len(w.childrenIds)+1] = child.id
		child.parent = w
	}

	return nil
}

func (w *Widget) GetChild(id string) *Widget {
	return w.children[id]
}

func (w *Widget) RemoveChild(id string) {
	child := w.children[id]
	if child == nil {
		return
	}

	delete(w.children, id)
	emitEvent("delete-widget", child.id)
}
