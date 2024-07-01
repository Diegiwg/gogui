package lib

import "fmt"

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
	w.emitContentUpdate()
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
	w.emitContentUpdate()
}
