package log

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/armnerd/go-skeleton/config"
	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
)

var mode int
var CONSOLE_MODE = 1
var FILE_MODE = 2

func init() {
	// 默认输出到终端，用 ELK 收集
	mode = CONSOLE_MODE
}

func getTraceId(c *gin.Context) string {
	traceId := c.GetString("traceId")
	return traceId
}

// Debug debug
func Debug(c *gin.Context, tag string, info interface{}) {
	logToFile(c, logrus.DebugLevel, tag, info)
}

// Info info
func Info(c *gin.Context, tag string, info interface{}) {
	logToFile(c, logrus.InfoLevel, tag, info)
}

// Warn warn
func Warn(c *gin.Context, tag string, info interface{}) {
	logToFile(c, logrus.WarnLevel, tag, info)
}

// Fatal fatal
func Fatal(c *gin.Context, tag string, info interface{}) {
	logToFile(c, logrus.FatalLevel, tag, info)
}

// Error error
func Error(c *gin.Context, tag string, info interface{}) {
	logToFile(c, logrus.ErrorLevel, tag, info)
}

// Panic panic
func Panic(c *gin.Context, tag string, info interface{}) {
	logToFile(c, logrus.PanicLevel, tag, info)
}

// 记录日志到文件
func logToFile(c *gin.Context, level logrus.Level, tag string, something interface{}) {
	var info string
	switch value := something.(type) {
	case string:
		info = value
	case map[string]interface{}:
		jsonStr, err := json.Marshal(value)
		if err != nil {
			return
		}
		info = string(jsonStr)
	default:
		return
	}
	var out io.Writer
	if mode == CONSOLE_MODE {
		// 输出终端
		out = os.Stdout
	} else if mode == FILE_MODE {
		// 日志文件
		logFilePath := config.AppRoot + os.Getenv("LOG_FILE")
		today := time.Now().Format("2006-01-02")
		logFileName := today + ".log"
		fileName := path.Join(logFilePath, logFileName)
		// 日志文件不存在时创建
		isExist, _ := pathExists(fileName)
		if !isExist {
			f, err := os.Create(fileName)
			if err != nil {
				return
			}
			defer f.Close()
		}
		var err error
		out, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Printf("Error while log: %v \n", err)
			return
		}
	}

	// 实例
	logger := logrus.New()
	logger.Out = out
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(logrus.DebugLevel)

	// 记录日志
	app := os.Getenv("APP_NAME")
	server := os.Getenv("SERVER_NUM")
	field := logrus.Fields{
		"app":     app,
		"server":  server,
		"tag":     tag,
		"traceId": getTraceId(c),
	}
	switch level {
	case logrus.DebugLevel:
		logger.WithFields(field).Debug(info)
	case logrus.InfoLevel:
		logger.WithFields(field).Info(info)
	case logrus.WarnLevel:
		logger.WithFields(field).Warn(info)
	case logrus.FatalLevel:
		logger.WithFields(field).Fatal(info)
	case logrus.ErrorLevel:
		logger.WithFields(field).Error(info)
	case logrus.PanicLevel:
		logger.WithFields(field).Panic(info)
	}
}

// 判断文件是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
