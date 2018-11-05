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
// @Summary 用户信息分页查询接口
// @Description 用户信息分页查询
// @Tags UserController
// @Accept json
// @Produce json
// @Param page query string true "页码"
// @Param page_size query string true "每页显示最大行"
// @Param username query string false "用户名"
// @Param phone query string false "电话"
// @Success 200 {object} helper.PageBean
// @Router /api/get_user_page [get]
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
// @Summary 用户信息删除接口
// @Description 删除用户信息
// @Tags UserController
// @Accept json
// @Produce json
// @Param id query string true "用户记录id"
// @Success 200 {object} helper.JsonObject
// @Router /api/delete_user [post]
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
// @Summary 获取所有用户接口
// @Description 获取所有用户信息
// @Tags UserController
// @Accept json
// @Produce json
// @Success 200 {object} helper.JsonObject
// @Router /api/get_all_users [get]
func GetAllUsers(context *gin.Context) {
	userService := service.UserServiceInstance(repositories.UserRepositoryInstance(helper.SQL))
	users := userService.GetAll()
	context.JSON(http.StatusOK, helper.JsonObject{
		Code:    "1",
		Content: users,
	})
	return
}
