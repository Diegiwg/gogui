package lib

import "fmt"

type WidgetKind int

const (
	WidgetElement = iota
	WidgetLabel
	WidgetButton
	WidgetGrid
)

func (wk *WidgetKind) String(w *Widget) string {
	switch *wk {
	case WidgetElement:
		return fmt.Sprintf("WidgetElement(%s)", w.GetData("tag").(string))
	case WidgetLabel:
		return "WidgetLabel"
	case WidgetButton:
		return "WidgetButton"
	case WidgetGrid:
		return fmt.Sprintf("WidgetGrid(%d, %d)", w.GetData("rows").(int), w.GetData("cols").(int))
	}
	return "WidgetKind"
}
