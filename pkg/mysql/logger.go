package mysql

import (
	"fmt"
	"time"

	syslog "github.com/armnerd/go-skeleton/pkg/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"

	"context"
)

type traceLogger struct {
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
	Colorful      bool
}

// 自定义日志
func NewTraceLogger(logLevel logger.LogLevel, slowThreshold time.Duration) *traceLogger {
	l := &traceLogger{}
	l.LogLevel = logLevel
	l.SlowThreshold = slowThreshold
	l.Colorful = true
	return l
}

// LogMode log mode
func (l *traceLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (l traceLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		output := append([]interface{}{utils.FileWithLineNum()}, data...)
		output = append(output, ctx.Value("traceId"))
		fmt.Printf(msg, output...)
	}
}

// Warn print warn messages
func (l traceLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		output := append([]interface{}{utils.FileWithLineNum()}, data...)
		output = append(output, ctx.Value("traceId"))
		fmt.Printf(msg, output...)
	}
}

// Error print error messages
func (l traceLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		output := append([]interface{}{utils.FileWithLineNum()}, data...)
		output = append(output, ctx.Value("traceId"))
		fmt.Printf(msg, output...)
	}
}

// Trace print sql message
func (l traceLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.LogLevel >= logger.Error:
			sql, rows := fc()
			c := &gin.Context{}
			c.Set("traceId", ctx.Value("traceId"))
			info := make(map[string]interface{})
			info["file"] = utils.FileWithLineNum()
			info["error"] = err
			info["time"] = float64(elapsed.Nanoseconds()) / 1e6
			info["sql"] = sql
			if rows == -1 {
				info["rows"] = "-"
				syslog.Error(c, "Sql Error", info)
			} else {
				info["rows"] = rows
				syslog.Error(c, "Sql Error", info)
			}
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
			sql, rows := fc()
			c := &gin.Context{}
			c.Set("traceId", ctx.Value("traceId"))
			info := make(map[string]interface{})
			info["file"] = utils.FileWithLineNum()
			info["error"] = fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
			info["time"] = float64(elapsed.Nanoseconds()) / 1e6
			info["sql"] = sql
			if rows == -1 {
				info["rows"] = "-"
				syslog.Warn(c, "Sql Slow", info)
			} else {
				info["rows"] = rows
				syslog.Warn(c, "Sql Slow", info)
			}
		case l.LogLevel >= logger.Info:
			sql, rows := fc()
			c := &gin.Context{}
			c.Set("traceId", ctx.Value("traceId"))
			info := make(map[string]interface{})
			info["file"] = utils.FileWithLineNum()
			info["error"] = ""
			info["time"] = float64(elapsed.Nanoseconds()) / 1e6
			info["sql"] = sql
			if rows == -1 {
				info["rows"] = "-"
				syslog.Info(c, "Sql Result", info)
			} else {
				info["rows"] = rows
				syslog.Info(c, "Sql Result", info)
			}
		}
	}
}
