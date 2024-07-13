package lib

type WidgetAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (w *Widget) Attributes() []WidgetAttribute {
	return w.attributes
}

func (w *Widget) SetAttribute(name, value string) {
	exists := w.HasAttribute(name)
	if exists {
		w.RemoveAttribute(name)
	}

	w.attributes = append(w.attributes, WidgetAttribute{name, value})
}

func (w *Widget) GetAttribute(name string) string {
	for _, attr := range w.attributes {
		if attr.Name == name {
			return attr.Value
		}
	}

	return ""
}

func (w *Widget) RemoveAttribute(name string) {
	var newAttributes []WidgetAttribute
	for _, attr := range w.attributes {
		if attr.Name == name {
			continue
		}
		newAttributes = append(newAttributes, attr)
	}
	w.attributes = newAttributes
}

func (w *Widget) HasAttribute(name string) bool {
	for _, attr := range w.attributes {
		if attr.Name == name {
			return true
		}
	}

	return false
}
