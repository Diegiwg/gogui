package gogui

import (
	"fmt"

	gogui_css "github.com/Diegiwg/gogui/gogui/css"
	gogui_js "github.com/Diegiwg/gogui/gogui/js"
)

func bundleMeta() string {
	return "<title>gogui</title>"
}

func bundleJs() string {
	data := "<script type=\"module\" defer>%s</script>"
	content := gogui_js.ButtonHandler
	return fmt.Sprintf(data, content)
}

func bundleCss() string {
	data := "<style>%s</style>"
	content := gogui_css.GridWidget
	return fmt.Sprintf(data, content)
}

func bundle() string {
	html := "<html><head>"

	html += bundleMeta()
	html += bundleJs()
	html += bundleCss()

	html += "</head><body>"

	return html
}
