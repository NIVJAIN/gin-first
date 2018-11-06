package model

import (
	"gin-first/helper"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// 功能菜单结构体
type Function struct {

	/** 主键 id */
	ID string `gorm:"type:varchar(36);primary_key" form:"id"`

	/** 功能名称 */
	FunName string

	/** 访问路径 */
	FunUrl string

	/** 权限功能等级 */
	funLevel int

	/** 是否生成菜单  */
	IsMenu bool

	/** 图标 */
	FunIcon string

	/** 序号 */
	Seq int

	/** 父功能 id */
	PId *string

	/** 父功能 */
	ParentFunction *Function `gorm:"foreignkey:PId;save_associations:false:" `

	/** 子功能 */
	ChildFunctions []*Function `gorm:"foreignkey:ID"`

	/** 对应的角色 */
	Roles []*Role `gorm:"many2many:role_functions;" json:"-"`

	/** 增删改的时间 */
	CrudTime
}

// 插入前生成主键
func (function *Function) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	if err != nil {
		helper.ErrorLogger.Errorln("生成UUID时发生异常: %s", err)
		return err
	}
	scope.Set("ID", &id)
	function.ID = id.String()
	return nil
}

func init() {
	// 创建或更新表结构
	helper.SQL.AutoMigrate(&Function{})
}
