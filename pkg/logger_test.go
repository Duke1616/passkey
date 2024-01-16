package pkg

import (
	"context"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log/slog"
	"os"
	"passkey-demo/pkg/logger"
	"testing"
)

func TestZapLogging(t *testing.T) {
	cfg := zap.NewDevelopmentConfig()
	err := viper.UnmarshalKey("log", &cfg)
	if err != nil {
		panic(err)
	}

	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	//ls := logger.NewZapLogger(l)

	l.Debug("你好", zap.Any("heelo", "123"))
	//ls.Debug("输入参数", logger.Field{Key: "username", Val: "passkey"})
}

func TestSLogger(t *testing.T) {
	opts := &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}
	l := slog.New(slog.NewTextHandler(os.Stdout, opts))

	ls := logger.NewSLogger(l)
	l.WithGroup("group").LogAttrs(context.Background(), slog.LevelInfo, "msg", slog.Int("a", 1), slog.Int("b", 2))

	ls.Error("你好")
	//l.LogAttrs(context.Background(), slog.LevelInfo, "msg", slog.Group("s", slog.Int("a", 1), slog.Int("b", 2)))

	//l.Info("hello")
	//ls.Debug("输入参数", logger.Field{Key: "username", Val: "passkey"})
}
