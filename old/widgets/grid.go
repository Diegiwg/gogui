package gogui_widgets

import (
	"fmt"
	"strconv"
)

func (w *Widget) Grid(rows int, cols int) *Widget {
	w.Kind = "Grid"
	w.SetData("rows", rows)
	w.SetData("cols", cols)
	w.SetData("children", make(map[int]*Widget))
	w.Html = w.gridHtml

	return w
}

func (w *Widget) gridHtml(id string) string {
	rows := w.GetData("rows").(int)
	cols := w.GetData("cols").(int)

	html := "<div "
	html += fmt.Sprintf("id=\"%s\" ", id)
	html += "class=\"grid-widget\" "
	html += " style=\""
	html += fmt.Sprintf("grid-template-rows: repeat(%d, 1fr); ", rows)
	html += fmt.Sprintf("grid-template-columns: repeat(%d, 1fr);", cols)
	html += "\">%s</div>"

	children := w.GetData("children").(map[int]*Widget)
	counter := len(children)

	content := ""
	for i := 1; i <= counter; i++ {
		content += children[i].Html(strconv.Itoa(i))
	}

	return fmt.Sprintf(html, content)
}
