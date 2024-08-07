package lib

import (
	"fmt"

	lib_css "github.com/Diegiwg/gogui/lib/css"
	lib_js "github.com/Diegiwg/gogui/lib/js"
)

func bundleMeta() string {
	return "<title>gogui</title>"
}

func bundleJs(app *App) string {
	data := "<script type=\"module\" defer>%s</script>"

	var content string
	content += lib_js.Base
	content += lib_js.Utils
	content += lib_js.DeleteWidget
	content += lib_js.RenderHtml
	content += fmt.Sprintf(lib_js.WsClient, app.config.serverAddress())

	return fmt.Sprintf(data, content)
}

func bundleCss() string {
	data := "<style>%s</style>"
	content := lib_css.GridWidget
	return fmt.Sprintf(data, content)
}

func bundle(app *App) string {
	html := "<html><head>"

	html += bundleMeta()
	html += bundleJs(app)
	html += bundleCss()

	html += "</head><body>"

	return html
}
