package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	filePath := "log/blog"
	linkName := "log/latest_log.log"
	// scr, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	// if err != nil {
	// 	fmt.Println("err:", err)
	// }
	logger := logrus.New()
	// logger.Out = scr //日志输出到文件

	logger.SetLevel(logrus.DebugLevel) //设置日志等级

	//日志文件按时间分割。使用file-rotatelogs和lfshook库
	logWriter, _ := rotatelogs.New(
		filePath+"%Y%m%d.log",
		rotatelogs.WithLinkName(linkName),         //设置软链接，链接到最新的日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     //最大保存时间：设置为7天
		rotatelogs.WithRotationTime(24*time.Hour), //多少时间分割一次：设置为24小时
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", //格式化输入日志的时间
	})

	logger.AddHook(Hook)

	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		stopTime := time.Since(startTime).Milliseconds()
		spendTime := fmt.Sprintf("%d ms", stopTime)
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknow"
		}
		statusCode := ctx.Writer.Status() //状态码
		clientIp := ctx.ClientIP()
		userAgent := ctx.Request.UserAgent() //客户端请求的浏览器
		dataSize := ctx.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := ctx.Request.Method
		path := ctx.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		if len(ctx.Errors) > 0 {
			entry.Error(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
