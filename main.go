package main

import (
	"gin-first/helper"
	"gin-first/routers"
	"gin-first/system"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func main() {
	ginConfig := system.GetGinConfig()
	gin.SetMode(ginConfig.RunMode)
	router := gin.New();
	router.Use(system.Logger(helper.AccessLogger),gin.Recovery())
	router.HandleMethodNotAllowed = ginConfig.HandleMethodNotAllowed
	router.Static("/page", "view");
	router.MaxMultipartMemory = ginConfig.MaxMultipartMemory
	routers.RegisterApiRoutes(router)
	routers.RegisterAppRoutes(router)
	routers.RegisterOpenRoutes(router)
	routers.RegisterAuthRoutes(router)
	serverConfig := system.GetServerConfig()
	server := &http.Server{
		Addr:           serverConfig.Addr,
		IdleTimeout:    serverConfig.IdleTimeout  * time.Second,
		ReadTimeout:    serverConfig.ReadTimeout  * time.Second,
		WriteTimeout:   serverConfig.WriteTimeout * time.Second,
		MaxHeaderBytes: serverConfig.MaxHeaderBytes,
		Handler:        router,
	}
	server.ListenAndServe();
}

func init()  {
	// 先读取服务配置文件
	err := system.LoadServerConfig("conf/server-config.yml");
	if err !=nil {
		helper.ErrorLogger.Errorln("读取服务配置错误：",err)
		os.Exit(3)
	}
}


