package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path"
	"time"
)

var Log *logrus.Logger

func LoggerMiddleware() gin.HandlerFunc {
	// 读取日志存储配置
	logFilePath := viper.GetString("logs.path")
	logFileName := viper.GetString("logs.name")
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	// 写入文件
	var (
		src *os.File
		err error
	)
	// 判断日志文件是否存在，存在直接打开，不存在则创建
	if _,err = os.Stat(fileName); os.IsNotExist(err){
		src, err = os.Create(fileName)
	}else{
		src, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}

	if err != nil {
		fmt.Println("err", err)
	}
	// 实例化
	Log = logrus.New()
	// 时间格式
	Log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})
	//设置日志级别
	Log.SetLevel(logrus.DebugLevel)
	//设置输出
	Log.Out = src


	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName + ".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	Log.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()
		//请求参数
		bodyBytes,_ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // write body back
		request := string(bodyBytes)
		//返回结果
		response,_ := c.Get("response")
		// 日志格式
		Log.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUrl,
			"req_body": request,
			"response": response,
		}).Info()
	}
}