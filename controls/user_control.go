package control

import (
	"gin-first/helper"
	"gin-first/models"
	"gin-first/repositories"
	"gin-first/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加用户接口
func SaveUser(context *gin.Context) {
	user := &model.User{}
	if err := context.Bind(user); err == nil {
		err = user.Validator()
		if err != nil {
			context.JSON(http.StatusOK,
				&helper.JsonObject{
					Code:    "0",
					Message: err.Error(),
				})
			return
		}
		user.DeletedAt = nil
		user.Role = nil
		user.RoleId = nil
		userService := service.UserServiceInstance(repositories.UserRepositoryInstance(helper.SQL))
		err := userService.SaveOrUpdate(user)
		if err == nil {
			context.JSON(http.StatusOK,
				&helper.JsonObject{
					Code:    "1",
					Message: helper.StatusText(helper.SaveStatusOK),
				})
			return
		} else {
			context.JSON(http.StatusOK,
				&helper.JsonObject{
					Code:    "0",
					Message: err.Error(),
				})
			return
		}
	} else {
		context.JSON(http.StatusUnprocessableEntity, helper.JsonObject{
			Code:    "0",
			Message: helper.StatusText(helper.BindModelErr),
			Content: err,
		})
	}
}

// 分页查询接口
func GetUserPage(context *gin.Context) {
	page, _ := strconv.Atoi(context.Query("page"))
	pageSize, _ := strconv.Atoi(context.Query("page_size"))
	username := context.Query("username")
	phone := context.Query("phone")
	userService := service.UserServiceInstance(repositories.UserRepositoryInstance(helper.SQL))
	pageBean := userService.GetPage(page, pageSize, &model.User{UserName: username, Phone: phone})
	context.JSON(http.StatusOK, helper.JsonObject{
		Code:    "1",
		Content: pageBean,
	})
}

// 删除接口
func DeleteUser(context *gin.Context) {
	id := context.Query("id")
	userService := service.UserServiceInstance(repositories.UserRepositoryInstance(helper.SQL))
	err := userService.DeleteByID(id)
	if err != nil {
		context.JSON(http.StatusOK, helper.JsonObject{
			Code:    "0",
			Message: helper.StatusText(helper.DeleteStatusErr),
			Content: err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, helper.JsonObject{
			Code:    "1",
			Message: helper.StatusText(helper.DeleteStatusOK),
		})
	}
	return
}

// 获取所有用户数据
func GetAllUsers(context *gin.Context) {
	userService := service.UserServiceInstance(repositories.UserRepositoryInstance(helper.SQL))
	users := userService.GetAll()
	context.JSON(http.StatusOK, helper.JsonObject{
		Code:    "1",
		Content: users,
	})
	return
}
