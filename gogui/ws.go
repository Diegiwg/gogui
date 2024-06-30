package gogui

import (
	"context"
	"encoding/json"
	"log"

	"nhooyr.io/websocket"
)

type Payload struct {
	Action string                 `json:"action"`
	Data   map[string]interface{} `json:"data"`
}

type WsClient struct {
	WsConn *websocket.Conn
	Ctx    *context.Context
}

func (c *WsClient) Send(data map[string]interface{}) {
	msg, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	c.WsConn.Write(*c.Ctx, websocket.MessageText, msg)
}

func wsHandler(httpCtx *HttpCtx) {
	w := *httpCtx.Response
	r := httpCtx.Request

	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.CloseNow()

	httpCtx.WsClient = &WsClient{WsConn: conn}

	for {
		ctx := context.Background()
		httpCtx.WsClient.Ctx = &ctx

		_, msg, err := conn.Read(ctx)
		if err != nil {
			break
		}

		var pl Payload
		err = json.Unmarshal(msg, &pl)
		if err != nil {
			log.Printf("%s", msg)
		}

		switch pl.Action {
		case "button-click":
			buttonHandler(httpCtx, pl.Data)
		}

		err = conn.Write(ctx, websocket.MessageText, make([]byte, 0))
		if err != nil {
			break
		}
	}
}

func buttonHandler(ctx *HttpCtx, data map[string]interface{}) {
	actionId, ok := data["id"].(string)
	if !ok {
		return
	}

	action, exists := ctx.App.actions["button-"+actionId]
	if !exists || action == nil {
		return
	}

	action(ctx, data)
}
