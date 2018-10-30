package control

import (
	"gin-first/helper"
	"gin-first/models"
	"gin-first/repositories"
	"gin-first/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 保存角色
func SaveRole(context *gin.Context) {
	var role model.Role
	if err := context.Bind(&role); err == nil {
		role.DeletedAt = nil
		roleService := service.RoleServiceInstance(repositories.RoleRepositoryInstance(helper.SQL))
		err := roleService.SaveOrUpdate(&role)
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
					Message: helper.StatusText(helper.SaveStatusErr),
					Content: err.Error(),
				})
			return
		}
	} else {
		context.JSON(http.StatusUnprocessableEntity, helper.JsonObject{
			Code:    "0",
			Message: helper.StatusText(helper.BindModelErr),
			Content: err,
		})
		return
	}
}
