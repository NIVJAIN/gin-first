package main

import (
	"gin-first/helper"
	"gin-first/routers"
	"gin-first/system"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

// @title gin-first API 文档
// @version 1.0
// @description 这是一个 gin 框架应用 swagger 的示例

// @contact.name YinYongTao
// @contact.url https://github.com/YinYongTao/gin-first
// @contact.email yongtao.yin@bsit.cn

// @licenes.name Apache 2.0
// @licenes.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 192.168.0.103:8080
// @BasePath /
func main() {
	ginConfig := system.GetGinConfig()
	gin.SetMode(ginConfig.RunMode)
	router := gin.New()
	router.Use(system.Logger(helper.AccessLogger), gin.Recovery())
	//配置跨域
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "ACCESS_TOKEN"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))
	router.HandleMethodNotAllowed = ginConfig.HandleMethodNotAllowed
	router.Static("/page", "view")
	router.MaxMultipartMemory = ginConfig.MaxMultipartMemory
	routers.RegisterApiRoutes(router)
	routers.RegisterAppRoutes(router)
	routers.RegisterOpenRoutes(router)
	routers.RegisterAuthRoutes(router)
	serverConfig := system.GetServerConfig()
	server := &http.Server{
		Addr:           serverConfig.Addr,
		IdleTimeout:    serverConfig.IdleTimeout * time.Second,
		ReadTimeout:    serverConfig.ReadTimeout * time.Second,
		WriteTimeout:   serverConfig.WriteTimeout * time.Second,
		MaxHeaderBytes: serverConfig.MaxHeaderBytes,
		Handler:        router,
	}
	server.ListenAndServe()
}

func init() {
	// 先读取服务配置文件
	err := system.LoadServerConfig("conf/server-config.yml")
	if err != nil {
		helper.ErrorLogger.Errorln("读取服务配置错误：", err)
		os.Exit(3)
	}
}
