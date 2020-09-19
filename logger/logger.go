package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//Logger instance
var (
	Log *zap.Logger
)

func init() {
	logConf := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	if Log, err = logConf.Build(); err != nil {
		panic(err)
	}
}

//Info log
func Info(msg string, tags ...zap.Field) {
	Log.Info(msg, tags...)
	Log.Sync()
}

//Error tag
func Error(msg string, err error, tags ...zap.Field) {
	if err != nil {
		tags = append(tags, zap.NamedError("error", err))
	}
	Log.Error(msg, tags...)
	Log.Sync()
}
