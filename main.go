package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/IanVzs/Snowflakes/pkgs/logging"
	"github.com/IanVzs/Snowflakes/pkgs/setting"
	_ "github.com/IanVzs/Snowflakes/pkgs/setting"
	_ "github.com/IanVzs/Snowflakes/pkgs/util"
	"github.com/IanVzs/Snowflakes/routers"
	_ "github.com/IanVzs/Snowflakes/services/ws_service"
)

func main() {

	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	s := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout * time.Second,
		WriteTimeout:   writeTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	s.ListenAndServe()
}
