package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		TimeKey:      "time",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	consoleSink := zapcore.AddSync(os.Stdout)
	core := zapcore.NewCore(consoleEncoder, consoleSink, zap.
		InfoLevel)
	logger := zap.New(core)
	sugar := logger.Sugar()
	sugar.Infow("A group of walrus emerges from the ocean",
		"animal", "walrus",
		"size", 10,
	)
}
