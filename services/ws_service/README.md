# 使用示例
`ws_service`的`init`中含有了`HManager`的`run`.
详见: [此处](https://github.com/IanVzs/Snowflakes/tree/master/services/ws_service/ws.go)

```go
ws "github.com/IanVzs/Snowflakes/services/ws_service"

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
```