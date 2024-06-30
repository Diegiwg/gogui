package gogui

import (
	"log"
	"net/http"
)

type HttpHandler func(ctx *HttpCtx, data map[string]interface{})

type HttpCtx struct {
	App      *App
	Html     string
	Request  *http.Request
	Response *http.ResponseWriter
	WsClient *WsClient
}

func rootHandler(ctx *HttpCtx) {
	rw := *ctx.Response
	html, css := ctx.App.widgetTree.Render()

	rw.Write([]byte(
		ctx.Html + html + "</body><style>" + css + "</style></html>", // fmt.Sprintf don't work here because css have '%' in it
	))
}

func (app *App) requestHandler(w http.ResponseWriter, r *http.Request) {
	ctx := &HttpCtx{
		App:      app,
		Html:     bundle(app),
		Request:  r,
		Response: &w,
	}

	if r.Method != "GET" {
		return
	}

	log.Printf("REQUEST: %s", r.URL.Path)

	if r.URL.Path == "/" {
		rootHandler(ctx)
		return
	}

	if r.URL.Path == "/ws" {
		wsHandler(ctx)
		return
	}
}
