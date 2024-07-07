package lib

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
)

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

		handleEvent(event, conn, &ctx)

		err = conn.Write(ctx, websocket.MessageText, []byte("{}"))
		if err != nil {
			break
		}
	}
}
