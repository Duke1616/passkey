package config

import (
	"passkey-demo/pkg/logger"
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

func LoadConfigFromToml() error {
	var err error
	conf, err = TryLoadFromDisk()

	if err != nil {
		L.Error("Failed to load configuration from disk: %v", logger.Error(err))
	}

	return nil
}

func LoadConfigFromEnv() error {
	return nil
}
