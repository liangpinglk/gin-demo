package tools

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Sugar *zap.SugaredLogger

func InitLog() {
	fileOutputW := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "data/log/test.log",
		MaxSize:    2, //,egabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	})
	consoleOutputW := zapcore.AddSync(os.Stdout)
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			fileOutputW,
			zap.InfoLevel,
		),
		zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			consoleOutputW,
			zap.InfoLevel),
	)
	logger := zap.New(core)
	sugar := logger.Sugar()
	Sugar = sugar
}
