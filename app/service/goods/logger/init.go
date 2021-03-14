package logger

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var logger *zap.Logger

func InitLogger() {
	var coreArr []zapcore.Core

	// 获取编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 指定时间格式
	// CapitalColorLevelEncoder: 按级别显示不同颜色; CapitalLevelEncoder: 不显示颜色
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder      	// 显示完整文件路径
	// NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// 日志级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // error
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // info 和 debug
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	// info文件writeSyncer
	infoViper := viper.Sub("log.info")
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   infoViper.GetString("file"),
		MaxSize:    infoViper.GetInt("max-size"),
		MaxBackups: infoViper.GetInt("backups"),
		MaxAge:     infoViper.GetInt("max-Age"),
		Compress:   infoViper.GetBool("compress"),
	})
	// 第三个及之后的参数为写入文件的日志级别, ErrorLevel模式只记录error级别的日志
	infoFileCore := zapcore.NewCore(encoder,
		zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority)
	// error文件writeSyncer
	errViper := viper.Sub("log.err")
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   errViper.GetString("file"),
		MaxSize:    errViper.GetInt("max-size"),
		MaxBackups: errViper.GetInt("backups"),
		MaxAge:     errViper.GetInt("max-Age"),
		Compress:   errViper.GetBool("compress"),
	})
	// 第三个及之后的参数为写入文件的日志级别, ErrorLevel模式只记录error级别的日志
	errorFileCore := zapcore.NewCore(encoder,
		zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority)
	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)
	logger = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller(), zap.AddCallerSkip(1))
	Info("日志初始化成功")
}

func Info(msg string) {
	logger.Info(msg)
}

func InfoDetail(msg string, err error) {
	logger.Info(msg, zap.String("details", fmt.Sprintf("%+v", err)))
}

func Warn(msg string) {
	logger.Warn(msg)
}

func WarnDetail(msg string, err error) {
	logger.Warn(msg, zap.String("details", fmt.Sprintf("%+v", err)))
}

func Debug(msg string) {
	logger.Debug(msg)
}

func DebugDetail(msg string, err error) {
	logger.Debug(msg, zap.String("details", fmt.Sprintf("%+v", err)))
}

func Error(msg string) {
	logger.Error(msg)
}

func ErrorDetail(msg string, err error) {
	logger.Error(msg, zap.String("details", fmt.Sprintf("%+v", err)))
}

func Panic(msg string) {
	logger.Panic(msg)
}

func PanicDetail(msg string, err error) {
	logger.Panic(msg, zap.String("details", fmt.Sprintf("%+v", err)))
}
