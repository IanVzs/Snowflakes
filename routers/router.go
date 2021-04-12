package routers

import (
	"github.com/gin-gonic/gin"

	_ "github.com/IanVzs/Snowflakes/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	test "github.com/IanVzs/Snowflakes/routers/api/test"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/test")
	// apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/get", test.GoTest)
	}

	return r
}
