package lib

type WidgetData map[string]interface{}

func (w *Widget) SetData(key string, value interface{}) {
	w.data[key] = value
	// w.emitContentUpdate()
}

func (w *Widget) GetData(key string) interface{} {
	return w.data[key]
}

func (w *Widget) HasData(key string) bool {
	_, ok := w.data[key]
	return ok
}

func (w *Widget) DeleteData(key string) {
	delete(w.data, key)
	// w.emitContentUpdate()
}

func (w *Widget) ClearData() {
	w.data = make(WidgetData)
	// w.emitContentUpdate()
}
