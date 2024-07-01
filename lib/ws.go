package lib

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

var wsConn *websocket.Conn
var wsCtx *context.Context

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.CloseNow()

	for {
		ctx := context.Background()

		_, msg, err := conn.Read(ctx)
		if err != nil {
			break
		}

		wsConn = conn
		wsCtx = &ctx

		var payload EventPayload
		err = json.Unmarshal(msg, &payload)
		if err != nil {
			log.Printf("%s", msg)
		}

		if payload.Id == "" {
			log.Println("empty id in event payload")
			continue
		}

		widget := dom.GetWidget(payload.Id)
		if widget == nil {
			log.Println("widget not found")
			continue
		}

		event := widget.GetEvent(payload.Action)
		if event == nil {
			log.Println("event not found")
			continue
		}

		(*event)(widget, &payload)

		err = conn.Write(ctx, websocket.MessageText, make([]byte, 0))
		if err != nil {
			break
		}
	}
}

func (widget *Widget) emitContentUpdate() {
	if wsConn == nil || wsCtx == nil {
		return
	}

	wsConn.Write(*wsCtx, websocket.MessageText, []byte("update-element-content:"+widget.id+"|"+widget.Html()))
}
