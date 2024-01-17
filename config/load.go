package config

import (
	"log/slog"
)

var (
	conf *Config
)

// C 全局配置对象
func C() *Config {
	if conf == nil {
		panic("Load Config first")
	}

	return conf
}

func LoadConfig() error {
	var err error
	conf, err = TryLoadFromDisk()

	if err != nil {
		slog.Error("Failed to load configuration from disk: ", err)
	}

	return nil
}
