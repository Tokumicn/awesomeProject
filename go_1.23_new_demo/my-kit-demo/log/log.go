package log

import (
	"log/slog"
	"os"
	"strings"
	"sync"
)

const (
	WriterTypeText = "Text"
	WriterTypeJson = "Json"

	StoreTypeFile    = "File"
	StoreTypeConsole = "Console"
)

var (
	SLog          *slog.Logger
	hotUpdateLock sync.Mutex // 热更新锁
)

// LoggerConfig 日志配置
type LoggerConfig struct {
	Level      int    // 日志级别
	WriterType string // 输出类型 Text、Json (Default)
	StoreType  string // 存储类型 File、Console (Default)
}

// NewLogger  新建Logger
//
//	writerType 输出类型 Text、JSON(Default)
//	storeType  存储类型 Console、File
func NewLogger(cfg LoggerConfig) *slog.Logger {

	var (
		logLevel              slog.Level
		level                 int    = cfg.Level
		writerType, storeType string = cfg.WriterType, cfg.StoreType
	)

	// 默认日志级别 Debug
	if level == 0 {
		logLevel = slog.LevelInfo
	}

	// 默认输出类型 Text
	if len(writerType) <= 0 {
		writerType = WriterTypeJson
	}
	writerType = strings.TrimSpace(writerType)

	// 默认存储类型 Console
	if len(storeType) <= 0 {
		storeType = StoreTypeConsole
	}
	storeType = strings.Trim(storeType, "\n")

	// TODO 输入到文件
	writer := os.Stderr

	loggerHandlerOpt := &slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
	}

	hotUpdateLock.Lock()
	if writerType == WriterTypeText {
		SLog = slog.New(slog.NewTextHandler(writer, loggerHandlerOpt))
		SLog.Warn("[slog.New] TextHandler")
	} else if writerType == WriterTypeJson {
		SLog = slog.New(slog.NewJSONHandler(writer, loggerHandlerOpt))
		SLog.Warn("[slog.New] JSONHandler")
	}
	hotUpdateLock.Unlock()

	return SLog
}

// SetLoggerLevel  设置日志级别
func SetLoggerLevel(level int) {

	var slogLevel slog.Level
	switch level {
	case int(slog.LevelDebug):
		slogLevel = slog.LevelDebug
	case int(slog.LevelInfo):
		slogLevel = slog.LevelInfo
	case int(slog.LevelWarn):
		slogLevel = slog.LevelWarn
	case int(slog.LevelError):
		slogLevel = slog.LevelError
	default:
		slogLevel = slog.LevelInfo
	}

	slog.SetLogLoggerLevel(slogLevel)
}
