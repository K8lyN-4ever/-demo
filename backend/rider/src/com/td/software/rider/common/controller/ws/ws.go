package ws_controller

import (
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	ws "rider/src/com/td/software/rider/common/util/websocket"
)

var store = sessions.NewCookieStore([]byte("token"))

func Client(context *gin.Context) {
	upGrande := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upGrande.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		//log.WSLog(fmt.Sprintf("websocket connect error: %s", context.Param("channel")))
		//format.NewResponseJson(context).Error(51001)
		return
	}
	var session, _ = store.Get(context.Request, "token")
	if session.Values["account"] != nil {
		//注册
		client := &ws.Client{
			Id:      session.Values["account"].(string),
			Socket:  conn,
			Message: make(chan []byte, 1024),
		}
		ws.WebsocketManager.RegisterClient(client)
		go client.Read()
		go client.Write()
	}

}
