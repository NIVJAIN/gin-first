package helper

// 定义一些系统常用的 错误码

const (

	BindModelErr          =  20200

	LoginStatusSQLErr      = 20319
	LoginStatusErr         = 20300
	LoginStatusOK          = 20301

	SaveStatusOK           = 20400
	SaveStatusErr          = 20401
	SaveObjIsNil           = 20402

	DeleteStatusOK         = 20403
	DeleteStatusErr        = 20404
	DeleteObjIsNil         = 20405

	UpdateObjIsNil         = 20406


	ExistSameNameErr       = 20501
	ExistSamePhoneErr      = 20502


)


var statusText = map[int]string{

	BindModelErr:          "模型封装异常！",

	LoginStatusSQLErr:     "用户登陆时更新登陆数据异常！",
	LoginStatusErr:        "用户名或密码错误!",
	LoginStatusOK:         "登陆成功！",

	SaveStatusOK:          "保存成功！",
	SaveStatusErr:         "保存失败！",
	SaveObjIsNil:          "保存的对象为空！",



	DeleteStatusOK:        "删除成功！",
	DeleteStatusErr:       "删除失败！",
	DeleteObjIsNil:        "删除的记录不存在！",

	UpdateObjIsNil:        "修改的记录不存在！",

	ExistSameNameErr:      "已存在同名记录！",
	ExistSamePhoneErr:     "已存在相同手机号！",



}

func StatusText(code int) string {
	return statusText[code]
}