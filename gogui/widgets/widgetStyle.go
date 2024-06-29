package gogui_widgets

import (
	"fmt"
)

type WidgetStyleData map[string]string

type WidgetStyle struct {
	Data            WidgetStyleData
	Width           string `default:""`
	Height          string `default:""`
	BackgroundColor string `default:""`
	TextColor       string `default:""`
}

func (s *WidgetStyle) toString(key string) string {
	switch key {
	case "width":
		return s.Width
	case "height":
		return s.Height
	case "background-color":
		return s.BackgroundColor
	case "text-color":
		return s.TextColor

	default:
		return ""
	}
}

func (s *WidgetStyle) parseData() string {
	if s.Data == nil {
		return ""
	}

	var style string = ""

	for key, value := range s.Data {
		style += fmt.Sprintf("%s: %s; ", key, value)
	}

	return style
}

func (s *WidgetStyle) String(id string) string {
	style := fmt.Sprintf("#%s { ", id)
	something := false

	if s.Width != "" {
		style += fmt.Sprintf("width: %s; ", s.toString("width"))
		something = true
	}

	if s.Height != "" {
		style += fmt.Sprintf("height: %s; ", s.toString("height"))
		something = true
	}

	if s.BackgroundColor != "" {
		style += fmt.Sprintf("background-color: %s; ", s.toString("background-color"))
		something = true
	}

	if s.TextColor != "" {
		style += fmt.Sprintf("color: %s; ", s.toString("text-color"))
		something = true
	}

	customStyle := s.parseData()
	if customStyle != "" {
		style += customStyle
		something = true
	}

	style += "}"

	if something {
		return style
	}

	return ""
}
