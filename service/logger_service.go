package service

import (
	"maker-checker/conf"
	"maker-checker/logger"
	"os"

	"go.uber.org/zap"
)

type LoggerService interface {
	GetInstance() *zap.Logger
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
}

type loggerService struct {
	logger *zap.Logger
}

func NewLoggerService(conf *conf.GbeConfig) LoggerService {
	lg := logger.InitLoggger(conf.Logger)
	if lg == nil {
		os.Exit(1)
	}
	lg.Info("Logger initialized")
	return &loggerService{
		logger: lg,
	}

}

func (log *loggerService) GetInstance() *zap.Logger {
	return log.logger
}

func (l *loggerService) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *loggerService) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *loggerService) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *loggerService) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}
