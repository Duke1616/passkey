package ioc

import (
	"github.com/Duke1616/passkey/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log/slog"
	"os"
)

func InitLogger() logger.Logger {
	cfg := zap.NewDevelopmentConfig()
	err := viper.UnmarshalKey("log", &cfg)
	if err != nil {
		panic(err)
	}
	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger.NewZapLogger(l)
}

func InitLoggerSlog() logger.Logger {
	opts := &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}
	l := slog.New(slog.NewTextHandler(os.Stdout, opts))
	return logger.NewSLogger(l)
}
