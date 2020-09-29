package config

import (
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger

func getLogWriter(logPath string, maxSize, maxAge, maxBackups int, compress bool) zapcore.WriteSyncer {
	fileName := filepath.Join(logPath, "bbs.log")
	lumberJackerLogger := &lumberjack.Logger{
		Filename:   fileName,   // log's path and name
		MaxSize:    maxSize,    // the maximum size of each file
		MaxAge:     maxAge,     // the maximum storage time of each file
		MaxBackups: maxBackups, //the maximum number of backup log files
		Compress:   compress,   // whether or not to compress
	}
	return zapcore.AddSync(lumberJackerLogger)
}

func getEncoder() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encodeConfig)
}

// InitLogger init a logger
func InitLogger(logPath string, maxSize, maxAge, maxBackups int, compress bool) {
	writeSyncer := getLogWriter(logPath, maxSize, maxAge, maxBackups, compress)
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	log = zap.New(core, zap.AddCaller()).Sugar()
}

// GetLogger return the logger to print log out of godis
func GetLogger() *zap.SugaredLogger {
	return log
}
