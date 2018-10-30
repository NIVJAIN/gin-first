package model

import (
	"gin-first/helper"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

// 定义 增删改时间 结构体
type CrudTime struct {

	/** 创建时间 */
	CreatedAt time.Time

	/** 更新时间 */
	UpdatedAt time.Time

	/** 删除时间 */
	DeletedAt *time.Time  `sql:"index"`
}


type User struct {

	/** 主键 */
	ID        string    `gorm:"type:varchar(36);primary_key"`

	/** 生日 */
	Birthday  time.Time `form:"birthday" time_format:"2006-01-02"`

	/** 年龄 */
	Age       *int

	/** 姓名 */
	UserName  string    `gorm:"type:varchar(32);unique_index;not null" form:"username"`

	/** 密码  */
	Password  string    `gorm:"type:varchar(64);not null" form:"password"`

	/** 电话 */
	Phone     string    `gorm:"type:varchar(11);unique" form:"phone"`

	/** 邮件 */
	Email     *string

	/** 登陆次数 */
	LogonCount int

	/** 状态  0 正常  */
	Status     int

	/** 最后一次登陆时间 */
	LoginTime  time.Time `gorm:"default:null"`

	/** 增删改的时间 */
	CrudTime

	/** 用户对应的角色 */
	Role    *Role         `gorm:"foreignkey:RoleId;save_associations:false:" `

	/** 外键 */
	RoleId  *string       `gorm:"type:varchar(36)"`

}

// 表结构初始化
func init()  {
	// 创建或更新表结构
	helper.SQL.AutoMigrate(&User{});
	// 生成外键约束
	helper.SQL.Model(&User{}).AddForeignKey("role_id","role(id)","no action","no action")
}

// 插入前生成主键
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	if err != nil {
		helper.ErrorLogger.Errorln("生成UUID时发生异常: %s", err)
		return err
	}
	scope.Set("ID", &id)
	user.ID =id.String()
	return nil
}


