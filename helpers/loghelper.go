package helper

import (
	"github.com/sirupsen/logrus"
	"os"
)

var ErrorLogger, AccessLogger, WorkLogger,SQLLogger *logrus.Logger

func Logger(outpath string) *logrus.Logger {
	logger := logrus.New();
	_, err := os.Stat(outpath);
	if os.IsNotExist(err) {
		// 文件不存在,创建
		os.Create(outpath)
	}
	file, err := os.OpenFile(outpath, os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		logger.Out = file;
	}else {
		logger.Info("打开 "+outpath+" 下的日志文件失败, 使用默认方式显示日志！")
	}
	return logger;
}

func init()  {
	WorkLogger   = Logger("logs/work.log");
	AccessLogger = Logger("logs/access.log");
	ErrorLogger  = Logger("logs/error.log");
	SQLLogger    = Logger("logs/sql.log")
}

