package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

// 定义 基于logrus 的log中间件
func Logger(log *logrus.Logger) gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.Request.URL.Path
		start :=time.Now()
		context.Next()
		stop :=time.Since(start)
		// 等待时间
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := context.Writer.Status()
		clientIP := context.ClientIP()
		clientUserAgent := context.Request.UserAgent()
		referer := context.Request.Referer()
		hostname, err := os.Hostname()
		if err !=nil {
			hostname = "unknow"
		}
		dataLength := context.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		entry := logrus.NewEntry(log).WithFields(logrus.Fields{
			"hostname":   hostname,
			"statusCode": statusCode,
			"latency":    latency, // time to process
			"clientIP":   clientIP,
			"method":     context.Request.Method,
			"path":       path,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
		})

		if len(context.Errors) > 0 {
			entry.Error(context.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("[%s] \"%s %s\" %d",  time.Now().Format("2006-01-02 15:04:05"), context.Request.Method, path, statusCode)
			if statusCode > 499 {
				entry.Error(msg)
			} else if statusCode > 399 {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}