package gogui_widgets

type Widget struct {
	Kind string
	Html func(id string) string
	Data map[string]interface{}
}

func NewWidget() *Widget {
	return &Widget{
		Data: make(map[string]interface{}),
	}
}

func (widget *Widget) SetData(key string, value interface{}) {
	widget.Data[key] = value
}

func (widget *Widget) GetData(key string) interface{} {
	return widget.Data[key]
}

func (widget *Widget) Child(newWidget *Widget) {
	switch widget.Kind {
	case "Grid":
		children := widget.GetData("children").(map[int]*Widget)
		children[len(children)+1] = newWidget

		widget.SetData("children", children)
	}
}
