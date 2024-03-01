package routers

import (
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/IanVzs/Snowflakes/docs"
	"github.com/IanVzs/Snowflakes/pkgs/logging"
	"github.com/IanVzs/Snowflakes/pkgs/setting"
	"github.com/IanVzs/Snowflakes/routers/api"
	test "github.com/IanVzs/Snowflakes/routers/api/test"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	if setting.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		gin.DisableConsoleColor()
		r.Use(logging.Ginzap(logging.AppLogger, time.RFC3339, true))
		r.Use(logging.RecoveryWithZap(logging.AppLogger, true))
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// test ws can use
	/*curl --include --no-buffer --header "Connection: Upgrade" --header "Upgrade: websocket" --header "Host: 127.0.0.1:8000"  --header "Sec-WebSocket-Key: zVM4LLeZBgoAzNyTtkEjxGVbUEk="  --header "Sec-WebSocket-Version: 13" http://127.0.0.1:8000/ws*/
	r.GET("/ws", WsEcho)
	r.GET("/ws_chat", Chat)
	apiv1 := r.Group("/api/v1")
	// apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", test.GoTest)
		apiv1.GET("/qa", api.QA)
	}

	return r
}
