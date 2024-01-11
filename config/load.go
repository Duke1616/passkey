package config

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

func LoadConfigFromToml(filePath ...string) error {
	return nil
}

func LoadConfigFromEnv() error {
	return nil
}
