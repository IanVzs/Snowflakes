package routers

import (
	"net/http"

	"github.com/IanVzs/Snowflakes/pkg/logging"
	ws "github.com/IanVzs/Snowflakes/service/ws_service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

// websocket echo demo
func WsEcho(c *gin.Context) {
	// 获取链接,允许跨域
	conn, error := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if error != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	defer conn.Close()

	// websocket connect
	for {
		mtye, message, err := conn.ReadMessage()
		if err != nil {
			logging.Errorf("ws conn.ReadMessage(mtye: %d): %s, ", err, mtye)
		}
		logging.Debugf("recv(%d): %s", mtye, message)
		err = conn.WriteMessage(mtye, message)
		if err != nil {
			logging.Errorf("ws conn.WriteMessage: %s", err)
			break
		}
	}
}

// WsPage is a websocket handler
func WsPage(c *gin.Context) {
	// change the reqest to websocket model
	conn, error := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if error != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	// websocket connect
	client := &ws.Client{ID: uuid.NewV4().String(), Socket: conn, Send: make(chan []byte)}

	logging.Infof("New Client Conn: %s", client.ID)
	ws.Manager.Register <- client

	go client.Read()
	go client.Write()
}
