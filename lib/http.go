package lib

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
)

var wsConn *websocket.Conn
var wsCtx *context.Context

func (app *App) requestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	log.Printf("REQUEST: %s", r.URL.Path)

	if r.URL.Path == "/" {
		w.Write([]byte(
			"<html>" + bundle(app) + "<body id=\"app\"></body></html>",
		))
		return
	}

	if r.URL.Path == "/ws" {
		wsHandler(w, r)
		return
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.CloseNow()

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
		defer cancel()

		_, data, err := conn.Read(ctx)
		if err != nil {
			break
		}

		var event Event
		err = json.Unmarshal(data, &event)
		if err != nil {
			break
		}

		wsConn = conn
		wsCtx = &ctx
		handleEvent(event)

		err = conn.Write(ctx, websocket.MessageText, []byte("{}"))
		if err != nil {
			break
		}
	}
}
