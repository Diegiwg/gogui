package lib

import (
	"log"
	"net/http"
)

func (app *App) requestHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		return
	}

	log.Printf("REQUEST: %s", r.URL.Path)

	if r.URL.Path == "/" {
		w.Write([]byte(
			"<html>" + bundle(app) + "<body>" + app.Root.Html() + "</body></html>",
		))
		return
	}

	if r.URL.Path == "/ws" {
		wsHandler(w, r)
		return
	}
}
