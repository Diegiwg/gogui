package gogui

import (
	"fmt"
	"net/http"
)

type HttpCtx struct {
	App      *App
	Html     string
	Request  *http.Request
	Response *http.ResponseWriter
}

func rootHandler(ctx *HttpCtx) {
	rw := *ctx.Response
	rw.Write([]byte(
		ctx.Html + ctx.App.widgetTree.Render() + "</body></html>", // fmt.Sprintf don't work here because css have '%' in it
	))
}

func buttonHandler(ctx *HttpCtx) {
	if !ctx.Request.URL.Query().Has("actionId") {
		return
	}

	actionId := ctx.Request.URL.Query().Get("actionId")

	if actionId == "" {
		return
	}

	action, exists := ctx.App.actions["button-"+actionId]
	if !exists {
		return
	}

	action(ctx)

	rw := *ctx.Response
	rw.Write([]byte("OK"))
}

func (a *App) requestHandler(w http.ResponseWriter, r *http.Request) {
	ctx := &HttpCtx{
		App:      a,
		Html:     bundle(),
		Request:  r,
		Response: &w,
	}

	if r.Method != "GET" {
		return
	}

	println(fmt.Sprintf("URL: %s", r.URL.Path))

	if r.URL.Path == "/" {
		rootHandler(ctx)
		return
	}

	if r.URL.Path == "/button" {
		buttonHandler(ctx)
		return
	}
}
