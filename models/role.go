package model

import (
	"gin-first/helper"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// 角色结构体
type Role struct {

	/** 主键id */
	ID           string      `gorm:"type:varchar(36);primary_key" form:"id"`

	/** 角色名称 */
	RoleName     string      `gorm:"type:varchar(32);unique;not null" form:"role_name"`

	/** 角色类别标识 */
	RoleKey      string      `gorm:"type:varchar(16);not null"  form:"role_key"`

	/** 角色描述 */
	Description  string      `gorm:"type:varchar(128)" form:"description"`

	/** 角色关联的功能 */
	Functions    []*Function `gorm:"many2many:role_functions;"`

	/** 增删改的时间 */
	CrudTime
}

// 插入前生成主键
func (role *Role) BeforeCreate(scope *gorm.Scope) error   {
	id, err := uuid.NewV4()
	if err != nil {
		helper.ErrorLogger.Errorln("生成UUID时发生异常: %s", err)
		return err
	}
	scope.Set("ID", &id)
	role.ID =id.String()
	return nil
}

func init()  {
	// 创建或更新表结构
	helper.SQL.AutoMigrate(&Role{});
}

