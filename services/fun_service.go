package service

import (
	"errors"
	"gin-first/helpers"
	"gin-first/models"
	"gin-first/repositories"
	"strings"
)

// function service 接口
type FunctionService interface {

	// 保存或修改
	SaveOrUpdate(function *model.Function) error

	// 根据id查询
	GetByID(id string) *model.Function

	// 根据功能权限名称查询
	GetByFunName(funname string) *model.Function

	// 根据 id 删除
	DeleteByID(id string) error

	// 查询所有
	GetAll() []*model.Function

	// 分页查询
	GetPage(page int, pageSize int, user *model.Function) *helper.PageBean
}

// function service 结构体
type functionService struct {

	/** 存储对象 */
	repo repositories.FunctionReposotory
}

func (fs *functionService) SaveOrUpdate(function *model.Function) error {
	if function == nil {
		return errors.New(helper.StatusText(helper.SaveObjIsNil))
	}
	// 判断 新增还是更新
	if strings.TrimSpace(function.ID) == "" {
		return fs.repo.Insert(function)
	} else {
		return fs.repo.Update(function)
	}
}

func (fs *functionService) GetByID(id string) *model.Function {
	if strings.TrimSpace(id) == "" {
		return nil
	}
	function := fs.repo.FindOne(id).(*model.Function)
	return function
}

func (fs *functionService) GetByFunName(funname string) *model.Function {
	function := fs.repo.FindSingle("fun_name = ?", funname).(*model.Function)
	return function
}

func (*functionService) DeleteByID(id string) error {
	panic("implement me")
}

func (*functionService) GetAll() []*model.Function {
	panic("implement me")
}

func (*functionService) GetPage(page int, pageSize int, user *model.Function) *helper.PageBean {
	panic("implement me")
}

var funServiceIns = &functionService{}

func FunctionServiceInstance(repo repositories.FunctionReposotory) FunctionService {
	funServiceIns.repo = repo
	return funServiceIns
}
