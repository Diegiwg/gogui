package lib

type WidgetAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (w *Widget) Attributes() []WidgetAttribute {
	return w.attributes
}
