package config

import "github.com/caarlos0/env/v6"

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
		return err
	}

	return nil
}

func LoadConfigFromEnv() error {
	conf = NewDefaultConfig()
	if err := env.Parse(conf); err != nil {
		return err
	}

	return nil
}
