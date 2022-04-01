package logging

import (
	"fmt"
	"os"
	"path"
	"seckill-jiujia/conf"
	"seckill-jiujia/pkg/file"
	"time"

	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	config *conf.Config
	log    *zap.Logger
	level  zapcore.Level
)

// Zap 日志初始化
func Init(c *conf.Config) {
	err := file.IsNotExistMkDir(c.LogInfo.Director)
	if err != nil {
		panic(err)
	}
	config = c
	switch c.LogInfo.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if (level == zap.DebugLevel || level == zap.ErrorLevel) && len(c.LogInfo.StacktraceKey) > 0 {
		log = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		log = zap.New(getEncoderCore())
	}

	if c.LogInfo.ShowLine {
		log.WithOptions(zap.AddCaller())
	}
}

// Debug 调试日志
func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

// Info 信息日志
func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

// Warn 警告日志
func Warn(msg string, fields ...zap.Field) {
	log.Warn(msg, fields...)
}

// Error 错误日志
func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

// Fatal 生成fatal日志
func Fatal(msg string, fields ...zap.Field) {
	log.Fatal(msg, fields...)
}

// Panic 异常日志
func Panic(msg string, fields ...zap.Field) {
	log.Panic(msg, fields...)
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (c zapcore.EncoderConfig) {
	c = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  config.LogInfo.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch config.LogInfo.EncodeLevel {
	case "LowercaseLevelEncoder": // 小写编码器(默认)
		c.EncodeLevel = zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder": // 小写编码器带颜色
		c.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder": // 大写编码器
		c.EncodeLevel = zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder": // 大写编码器带颜色
		c.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		c.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return c
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if config.LogInfo.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() (core zapcore.Core) {
	writer, err := GetWriteSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(config.LogInfo.Prefix + "2006/01/02 - 15:04:05.000"))
}

// GetWriteSyncer zap logger中加入file-rotatelogs
func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(config.LogInfo.Director, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName(config.LogInfo.LinkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if config.LogInfo.LogConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
