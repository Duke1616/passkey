package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"strings"
	"sync"
)

var (
	// singleton instance of config package
	_config = defaultConfig()
)

const (
	defaultConfigurationName = "config"
	defaultConfigurationPath = "config"
)

type viperConfig struct {
	cfg         *Config
	cfgChangeCh chan Config
	watchOnce   sync.Once
	loadOnce    sync.Once
}

func defaultConfig() *viperConfig {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.SetConfigType("toml")
	viper.SetConfigName(defaultConfigurationName)
	viper.AddConfigPath(path + "/" + defaultConfigurationPath)

	// Load from current working directory, only used for debugging
	viper.AddConfigPath(".")

	// Load from Environment variables
	viper.SetEnvPrefix("passkey")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return &viperConfig{
		cfg:         NewDefaultConfig(),
		cfgChangeCh: make(chan Config),
		watchOnce:   sync.Once{},
		loadOnce:    sync.Once{},
	}
}

func TryLoadFromDisk() (*Config, error) {
	return _config.loadFromDisk()
}

func (c *viperConfig) loadFromDisk() (*Config, error) {
	var err error
	c.loadOnce.Do(func() {
		//if err = viper.ReadInConfig(); err != nil {
		//	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		//		// 配置文件不存在，不影响从环境变量中获取
		//		slog.Warn("Config file not found: ", err.Error())
		//	} else {
		//		return
		//	}
		//}

		if err = viper.ReadInConfig(); err != nil {
			return
		}

		if err = viper.Unmarshal(c.cfg); err != nil {
			return
		}

		fmt.Print(c.cfg)
	})

	return c.cfg, err
}

// WatchConfigChange return config change channel
func WatchConfigChange() <-chan Config {
	return _config.watchConfig()
}

func (c *viperConfig) watchConfig() <-chan Config {
	c.watchOnce.Do(func() {
		viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			cfg := NewDefaultConfig()
			if err := viper.Unmarshal(cfg); err != nil {
				slog.Warn("config reload error: ", err)
			} else {
				c.cfgChangeCh <- *cfg
				conf = cfg
			}
		})
	})
	return c.cfgChangeCh
}
