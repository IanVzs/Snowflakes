package routers

import (
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/IanVzs/Snowflakes/docs"
	"github.com/IanVzs/Snowflakes/pkg/logging"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	test "github.com/IanVzs/Snowflakes/routers/api/test"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	r.Use(logging.Ginzap(logging.Logger, time.RFC3339, true), logging.RecoveryWithZap(logging.Logger, true))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// test ws can use
	/*curl --include --no-buffer --header "Connection: Upgrade" --header "Upgrade: websocket" --header "Host: 127.0.0.1:8000"  --header "Sec-WebSocket-Key: zVM4LLeZBgoAzNyTtkEjxGVbUEk="  --header "Sec-WebSocket-Version: 13" http://127.0.0.1:8000/ws*/
	r.GET("/ws", WsEcho)
	r.GET("/wspage", WsPage)
	apiv1 := r.Group("/api/test")
	// apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/get", test.GoTest)
	}

	return r
}
