package lib

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

type wsClient struct {
	wsConn *websocket.Conn
	wsCtx  *context.Context
}

var clients = map[string]*wsClient{}

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

		var clientId string
		for {
			clientId = generateRandomId(25)

			if _, ok := clients[clientId]; !ok {
				break
			}
		}

		// TODO: remove dead clients
		clients[clientId] = &wsClient{wsConn: conn, wsCtx: &ctx}

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
	if len(clients) == 0 {
		return
	}

	for _, client := range clients {
		if client.wsConn == nil || client.wsCtx == nil {
			continue
		}

		client.wsConn.Write(*client.wsCtx, websocket.MessageText, []byte("update-element-content:"+widget.id+"|"+widget.Html()))
	}

}
