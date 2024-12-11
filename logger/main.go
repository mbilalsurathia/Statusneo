package logger

import (
	"maker-checker/conf"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLoggger(conf conf.Logger) *zap.Logger {
	write := setWrite(conf.Path)

	level := setLevel(conf.Level)

	encoder := setEncoder()

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	var core zapcore.Core
	if conf.Mode == "debug" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, write, level)
	}

	caller := zap.AddCaller()

	development := zap.Development()

	filed := zap.Fields(zap.String("service", conf.Name))

	lg := zap.New(core, caller, development, filed)

	zap.ReplaceGlobals(lg)

	return lg
}

func setWrite(path string) zapcore.WriteSyncer {

	hook := lumberjack.Logger{
		Filename:   path,
		MaxSize:    10,
		MaxBackups: 30,
		MaxAge:     7,
		Compress:   true,
	}
	write := zapcore.AddSync(&hook)

	return write
}

func setLevel(l string) zapcore.Level {
	var level zapcore.Level
	switch l {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	return level
}

func setEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encoderConfig)

	return encoder
}
