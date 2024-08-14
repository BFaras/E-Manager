package logger
import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
	"os"
)

var zapLog *zap.Logger

func init() {
    var err error
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(os.Stderr),
		zapcore.DebugLevel,
	)
	zapLog = zap.New(core)

    if err != nil {
        panic(err)
    }
}

func Info(message string, fields ...zap.Field) {
    zapLog.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
    zapLog.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
    zapLog.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
    zapLog.Fatal(message, fields...)
}