// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package logger

import (
	"fmt"
	"github.com/ysicing/ergo/utils"
	"github.com/ysicing/go-utils/extime"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var (
	loger *zap.SugaredLogger
)

const (
	LogSize     = 100
	LogBackup   = 3
	LogAge      = 60
	LogCompress = true
)

func InitLogger() {
	encoder := getEncoder()
	writeSyncer := getLogWriter()
	if utils.DebugMode() {
		core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), zapcore.DebugLevel)
		//loger = zap.New(core, zap.AddCaller()).Sugar()
		loger = zap.New(core).Sugar()
	} else {
		core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout)), zapcore.DebugLevel)
		//loger = zap.New(core, zap.AddCaller()).Sugar()
		loger = zap.New(core).Sugar()
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = timeEncoder                       //zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func getLogWriter() zapcore.WriteSyncer {
	var logpath string
	if utils.RunLinux() {
		logpath = fmt.Sprintf("/var/log/ergo.debug.%v.log", extime.GetToday())
	} else {
		logpath = fmt.Sprintf("/tmp/ergo.debug.%v.log", extime.GetToday())
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logpath,
		MaxSize:    LogSize,
		MaxBackups: LogBackup,
		MaxAge:     LogAge,
		Compress:   LogCompress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Debug(msg string) {
	loger.Debug(msg)
}

func Info(msg string) {
	loger.Info(msg)
}

func Warn(msg string) {
	loger.Warn(msg)
}

func Error(msg string) {
	loger.Error(msg)
}

func Exit(msg string) {
	loger.Error(msg)
	os.Exit(0)
}
