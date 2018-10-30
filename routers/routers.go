package routers

import (
	"gin-first/controls"
	"gin-first/system"
	"github.com/gin-gonic/gin"
	"net/http"
)

// api 路由注册
func RegisterApiRoutes(router *gin.Engine) {
	api := router.Group("api")
	// 鉴权
	api.Use(system.JWTAuth())
	api.POST("save_user",control.SaveUser)
	api.GET("get_user_page",control.GetUserPage)
	api.POST("delete_user",control.DeleteUser)
	api.GET("get_all_users",control.GetAllUsers)

	api.POST("save_role",control.SaveRole)

}

// app 路由注册
func RegisterAppRoutes(router *gin.Engine) {
	app := router.Group("/app");
	// 鉴权
	app.Use(system.JWTAuth())
	app.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello APP");
	})

}

// 注册其他需要鉴权的接口
func RegisterAuthRoutes(router *gin.Engine)  {
	router.Use(system.JWTAuth())

}

// 注册不需要鉴权的 接口
func RegisterOpenRoutes(router *gin.Engine)  {
	router.POST("login", control.Login);
}
