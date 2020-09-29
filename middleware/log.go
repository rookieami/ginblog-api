package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

//日志中间件
func Logger() gin.HandlerFunc {
	filepath := "log/log"
	linkName := "latest_log.log" //软连接,指向最新日志文件
	scr, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println("err:", err)
	}
	logger := logrus.New()

	logger.Out = scr

	logger.SetLevel(logrus.DebugLevel)
	//日志分割
	logwrite, _ := retalog.New(
		filepath+"%Y%m%d.log",

		retalog.WithMaxAge(7*24*time.Hour),     //只保存一周
		retalog.WithRotationTime(24*time.Hour), //每天分割一次
		retalog.WithLinkName(linkName),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logwrite,
		logrus.FatalLevel: logwrite,
		logrus.DebugLevel: logwrite,
		logrus.WarnLevel:  logwrite,
		logrus.ErrorLevel: logwrite,
		logrus.PanicLevel: logwrite,
	}
	//实例化
	hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(hook)
	return func(c *gin.Context) {
		startTime := time.Now() //开始时间
		c.Next()
		stopTime := time.Since(startTime)                                                            //结束时间
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0)))) //整数
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknow"
		}
		statusCode := c.Writer.Status()   //状态码
		clientIP := c.ClientIP()          //用户ip
		userAget := c.Request.UserAgent() //客户端信息
		dataSize := c.Writer.Size()       //请求文件长度
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method //请求方法
		path := c.Request.RequestURI
		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"Status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIP,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAget,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypeAny).String())
		} else if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}

	}
}
