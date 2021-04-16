package routers

import (
	"net/http"

	"github.com/IanVzs/Snowflakes/pkgs/logging"
	ws "github.com/IanVzs/Snowflakes/services/ws_service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

// WsPage is a websocket handler
func WsEcho(c *gin.Context) {
	// change the reqest to websocket model
	conn, error := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if error != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	// websocket connect
	client := &ws.EchoClient{ID: uuid.NewV4().String(), Socket: conn, Send: make(chan []byte)}

	logging.Infof("New Client Conn: %s", client.ID)
	ws.Manager.Register <- client

	go client.Read()
	go client.Write()
	logging.Infof("New Client Done: %s", client.ID)
}

// Chat is a websocket chat handler
func Chat(c *gin.Context) {
	var hub *ws.Hub
	var ok bool
	chatID := c.Query("chat_id")
	if chatID != "" {
		conn, error := (&websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
		if error != nil {
			http.NotFound(c.Writer, c.Request)
			return
		}
		if hub, ok = ws.HManager.Hubs[chatID]; !ok {
			hub = ws.NewHub(chatID)
			logging.Infof("first in chatID： %s", chatID)
		} else {
			logging.Infof("join in chatID： %s", chatID)
		}
		ws.ServeWs(hub, conn)
	} else {
		return
	}
}
