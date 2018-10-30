package helper

import (
	"gin-first/system"
	"github.com/jinzhu/gorm"
	 _"github.com/jinzhu/gorm/dialects/mysql"
)

var SQL *gorm.DB

// 初始化连接
func init()  {
	// 先读取Token配置文件
	err := system.LoadDatasourceConfig("./conf/datasource.yml");
	if err !=nil {
		ErrorLogger.Errorln("读取数据库配置错误：",err);
	}
	datasource := system.GetDatasource();

	SQL, err = gorm.Open(datasource.Driver, datasource.Username+":"+datasource.Password+datasource.Url);
	if err !=nil {
		ErrorLogger.Errorln("连接数据库失败：",err);
	}
	SQL.DB().SetMaxOpenConns(datasource.MaxOpenConns);
	SQL.DB().SetMaxIdleConns(datasource.MaxIdleConns);
	//SQL.SetLogger(SQLLogger)
	SQL.LogMode(datasource.ShowSql);
	SQL.SingularTable(datasource.SingularTable);
}

