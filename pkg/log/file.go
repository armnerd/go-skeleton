package log

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/armnerd/go-skeleton/config"

	"github.com/sirupsen/logrus"
)

// Debug debug
func Debug(tag string, info interface{}) {
	logToFile(logrus.DebugLevel, tag, info)
	return
}

// Info info
func Info(tag string, info interface{}) {
	logToFile(logrus.InfoLevel, tag, info)
	return
}

// Warn warn
func Warn(tag string, info interface{}) {
	logToFile(logrus.WarnLevel, tag, info)
	return
}

// Fatal fatal
func Fatal(tag string, info interface{}) {
	logToFile(logrus.FatalLevel, tag, info)
	return
}

// Error error
func Error(tag string, info interface{}) {
	logToFile(logrus.ErrorLevel, tag, info)
	return
}

// Panic panic
func Panic(tag string, info interface{}) {
	logToFile(logrus.PanicLevel, tag, info)
	return
}

// 记录日志到文件
func logToFile(level logrus.Level, tag string, something interface{}) {
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
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Printf("Error while log: %v \n", err)
		return
	}

	// 实例
	logger := logrus.New()
	logger.Out = src
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(logrus.DebugLevel)

	// 记录日志
	app := os.Getenv("APP_NAME")
	server := os.Getenv("SERVER_NUM")
	field := logrus.Fields{
		"app":    app,
		"server": server,
		"tag":    tag,
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
	return
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
