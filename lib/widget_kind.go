package lib

import "fmt"

type WidgetKind int

const (
	WidgetElement WidgetKind = iota
	WidgetLabel
	WidgetButton
	WidgetGrid
	WidgetInput
)

func (wk *WidgetKind) String(w *Widget) string {
	switch *wk {
	case WidgetElement:
		return fmt.Sprintf("WidgetElement(%s)", w.GetData("tag").(string))
	case WidgetLabel:
		return "WidgetLabel"
	case WidgetButton:
		return "WidgetButton"
	case WidgetInput:
		return "WidgetInput"
	case WidgetGrid:
		return fmt.Sprintf("WidgetGrid(%d, %d)", w.GetData("rows").(int), w.GetData("cols").(int))
	}
	return "WidgetKind"
}
