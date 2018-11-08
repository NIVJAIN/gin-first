package control

import (
	"gin-first/helper"
	"gin-first/models"
	"gin-first/repositories"
	"gin-first/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// 添加、修改用户信息
// @Summary 添加、修改用户信息
// @Tags UserController
// @Accept json
// @Produce json
// @Param id             query string false "用户记录id,新增时id为空"
// @Param username       query string true  "用户名"
// @Param password       query string true  "密码"
// @Param phone          query string true  "电话号码"
// @Param email          query string true  "邮件"
// @Param merchant_no    query string true  "商户号"
// @Param role_id        query string true  "角色id"
// @Success 200 {object} helper.JsonObject
// @Router /api/save_user [post]
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

// 用户信息分页查询
// @Summary 用户信息分页查询
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

// 删除用户信息
// @Summary 删除用户信息
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

// 获取所有用户信息
// @Summary 获取所有用户信息
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

// 用户信息回显
// @Summary 用户信息回显,传id按id查，传username按用户名查
// @Tags UserController
// @Accept json
// @Produce json
// @Param id        query string false "用户记录id"
// @Param username  query string false "用户名"
// @Success 200 {object} model.User
// @Router /api/get_user [get]
func GetUser(context *gin.Context) {
	id := context.Query("id")
	username := context.Query("username")
	if strings.TrimSpace(username+id) == "" {
		context.JSON(http.StatusOK, helper.JsonObject{
			Code:    "0",
			Message: helper.StatusText(helper.NoneParamErr),
		})
	}
	userService := service.UserServiceInstance(repositories.UserRepositoryInstance(helper.SQL))
	var user *model.User
	if id != "" {
		user = userService.GetByID(id)
	}else {
		user = userService.GetByUserName(username)
	}
	context.JSON(http.StatusOK, helper.JsonObject{
		Code:    "1",
		Content: user,
	})
	return
}
