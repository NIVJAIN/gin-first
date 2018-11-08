package service

import (
	"errors"
	"gin-first/helpers"
	"gin-first/models"
	"gin-first/repositories"
	"strings"
)

// user_service 接口
type UserService interface {

	/** 保存或修改 */
	SaveOrUpdate(user *model.User) error

	/** 根据 id 查询 */
	GetByID(id string) *model.User

	/** 根据用户名查询 */
	GetByUserName(username string) *model.User

	// 根据电话号码查询
	GetByPhone(phone string) *model.User

	/** 根据 id 删除 */
	DeleteByID(id string) error

	/** 查询所有  */
	GetAll() []*model.User

	/** 分页查询 */
	GetPage(page int, pageSize int, user *model.User) *helper.PageBean
}

var userServiceIns = &userService{}

// 获取 userService 实例
func UserServiceInstance(repo repositories.UserRepository) UserService {
	userServiceIns.repo = repo
	return userServiceIns
}

// 结构体
type userService struct {

	/** 存储对象 */
	repo repositories.UserRepository
}

func (us *userService) GetByUserName(username string) *model.User {
	user := us.repo.FindSingle("user_name = ?", username)
	if user != nil {
		return user.(*model.User)
	}
	return nil
}

func (us *userService) GetByPhone(phone string) *model.User {
	user := us.repo.FindSingle("phone = ?", phone)
	if user != nil {
		return user.(*model.User)
	}
	return nil
}

func (us *userService) SaveOrUpdate(user *model.User) error {
	if user == nil {
		return errors.New(helper.StatusText(helper.SaveObjIsNil))
	}
	// 校验用户名是否重复
	userByName := us.GetByUserName(user.UserName)

	// 校验手机号码是否重复
	userByPhone := us.GetByPhone(user.Phone)
	if user.ID == "" {
		// 添加
		if userByName != nil && userByName.ID != "" {
			return errors.New(helper.StatusText(helper.ExistSameNameErr))
		}
		if userByPhone != nil && userByPhone.ID != "" {
			return errors.New(helper.StatusText(helper.ExistSamePhoneErr))
		}
		user.Password = helper.SHA256(user.Password)
		return us.repo.Insert(user)
	} else {
		// 修改
		persist := us.GetByID(user.ID)
		if persist == nil || persist.ID == "" {
			return errors.New(helper.StatusText(helper.UpdateObjIsNil))
		}
		if userByName != nil && userByName.ID != user.ID {
			return errors.New(helper.StatusText(helper.ExistSameNameErr))
		}

		if userByPhone != nil && userByPhone.ID != user.ID {
			return errors.New(helper.StatusText(helper.ExistSamePhoneErr))
		}
		user.Password = persist.Password
		return us.repo.Update(user)
	}
	return nil
}

func (us *userService) GetAll() []*model.User {
	users := us.repo.FindMore("1=1").([]*model.User)
	return users
}

func (us *userService) GetByID(id string) *model.User {
	if strings.TrimSpace(id) == "" {
		return nil
	}
	user := us.repo.FindOne(id).(*model.User)
	return user
}

func (us *userService) DeleteByID(id string) error {
	user := us.repo.FindOne(id).(*model.User)
	if user == nil || user.ID == "" {
		return errors.New(helper.StatusText(helper.DeleteObjIsNil))
	}
	err := us.repo.Delete(user)
	return err
}

func (us *userService) GetPage(page int, pageSize int, user *model.User) *helper.PageBean {
	andCons := make(map[string]interface{})
	if user != nil && user.UserName != "" {
		andCons["user_name LIKE ?"] = user.UserName + "%"
	}
	if user != nil && user.Phone != "" {
		andCons["phone LIKE ?"] = user.Phone + "%"
	}
	pageBean := us.repo.FindPage(page, pageSize, andCons, nil)
	return pageBean
}
