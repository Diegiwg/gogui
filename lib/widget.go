package lib

var widgetPool = make(map[string]*Widget, 0)

type WidgetTree map[string]*Widget
type WidgetTreeIds map[int]string
type WidgetRender func(obj *RenderHtmlPayload)

type Widget struct {
	id         string
	kind       WidgetKind
	render     WidgetRender
	data       WidgetData
	attributes []WidgetAttribute
	events     []WidgetEvent
	style      WidgetStyle

	parent      *Widget
	children    WidgetTree
	childrenIds WidgetTreeIds
}

func (w *Widget) GetId() string {
	return w.id
}

func newWidget() *Widget {
	id := generateId(10)
	w := &Widget{
		id:          id,
		data:        make(WidgetData),
		children:    make(WidgetTree),
		childrenIds: make(WidgetTreeIds),
		attributes:  make([]WidgetAttribute, 0),
		events:      make([]WidgetEvent, 0),
		style:       NewWidgetStyle(),
	}

	widgetPool[id] = w
	return w
}

func (w *Widget) Delete() {
	w.parent.RemoveChild(w.id)
}

func (w *Widget) Render() *RenderHtmlPayload {
	obj := &RenderHtmlPayload{
		Id:         w.id,
		Tag:        w.GetData("tag").(string),
		Attributes: w.Attributes(),
		Events:     w.Events(),
		Content:    w.Content(),
		Style:      w.Styles(),
	}

	if w.render != nil {
		w.render(obj)
	}

	obj.Children = make([]*RenderHtmlPayload, 0, len(w.children))
	for i := 0; i <= len(w.childrenIds); i++ {

		childId := w.childrenIds[i]
		if childId == "" {
			continue
		}

		child := w.children[childId]
		if child == nil {
			continue
		}

		obj.Children = append(obj.Children, child.Render())
	}

	return obj
}

func (w *Widget) Update() {
	emitRenderHtmlEvent(w.id, w)
}
