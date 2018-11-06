package control

import (
	"gin-first/helper"
	"gin-first/models"
	"gin-first/repositories"
	"gin-first/services"
	"gin-first/system"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// 用户登陆接口
// @Summary 用户登陆接口
// @Tags LoginController
// @Accept json
// @Produce json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {object} helper.JsonObject
// @Router /login [post]
func Login(context *gin.Context) {
	username := context.Query("username")
	password := context.Query("password")
	userService := service.UserServiceInstance(repositories.UserRepositoryInstance(helper.SQL))
	user := userService.GetByUserName(username)
	if user != nil && user.Password == helper.SHA256(password) {
		user.LogonCount += 1
		user.LoginTime = time.Now()
		err := userService.SaveOrUpdate(user)
		if err == nil {
			generateToken(context, user)
		} else {
			context.JSON(http.StatusOK, helper.JsonObject{
				Code:    "0",
				Message: helper.StatusText(helper.LoginStatusSQLErr),
				Content: err,
			})
		}
	} else {
		context.JSON(http.StatusOK, helper.JsonObject{
			Code:    "0",
			Message: helper.StatusText(helper.LoginStatusErr),
		})
	}
}

// 生成令牌
func generateToken(context *gin.Context, user *model.User) {
	j := system.NewJWT()
	claims := system.CustomClaims{
		ID:    user.ID,
		Name:  user.UserName,
		Phone: user.Phone,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() + system.GetTokenConfig().ActiveTime),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + system.GetTokenConfig().ExpiredTime*3600), // 过期时间
			Issuer:    system.GetTokenConfig().Issuer,
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		context.JSON(http.StatusOK, helper.JsonObject{
			Code:    "0",
			Message: err.Error(),
		})
		context.Abort()
	}
	context.JSON(http.StatusOK, helper.JsonObject{
		Code:    "0",
		Message: helper.StatusText(helper.LoginStatusOK),
		Content: gin.H{"ACCESS_TOKEN": token},
	})
}

func init() {
	// 先读取Token配置文件
	err := system.LoadTokenConfig("./conf/token-config.yml")
	if err != nil {
		helper.ErrorLogger.Errorln("读取Token配置错误：", err)
	}
	if len(strings.TrimSpace(system.GetTokenConfig().SignKey)) > 0 {
		system.SetSignKey(system.GetTokenConfig().SignKey)
	}
}
