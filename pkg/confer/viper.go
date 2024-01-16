package confer

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"passkey-demo/config"
	"passkey-demo/pkg/logger"
	"strings"
	"sync"
)

var (
	// singleton instance of config package
	_config = defaultConfig()
	L       logger.Logger
)

const (
	defaultConfigurationName = "config"
	defaultConfigurationPath = "/config"
)

type viperConfig struct {
	cfg         *config.Config
	cfgChangeCh chan config.Config
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
	viper.AddConfigPath(path + defaultConfigurationPath)

	// Load from current working directory, only used for debugging
	viper.AddConfigPath(".")

	// Load from Environment variables
	viper.SetEnvPrefix("passkey")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return &viperConfig{
		cfg:         config.NewConfig(),
		cfgChangeCh: make(chan config.Config),
		watchOnce:   sync.Once{},
		loadOnce:    sync.Once{},
	}
}

func TryLoadFromDisk() (*config.Config, error) {
	return _config.loadFromDisk()
}

func (c *viperConfig) loadFromDisk() (*config.Config, error) {
	var err error
	c.loadOnce.Do(func() {
		if err = viper.ReadInConfig(); err != nil {
			return
		}

		if err = viper.Unmarshal(c.cfg); err != nil {
			return
		}
	})

	return c.cfg, err
}

// WatchConfigChange return config change channel
func WatchConfigChange() <-chan config.Config {
	return _config.watchConfig()
}

func (c *viperConfig) watchConfig() <-chan config.Config {
	c.watchOnce.Do(func() {
		viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			cfg := config.NewConfig()
			if err := viper.Unmarshal(cfg); err != nil {
				L.Warn("config reload error: %v", logger.Error(err))
			} else {
				c.cfgChangeCh <- *cfg
				conf = cfg
			}
		})
	})
	return c.cfgChangeCh
}
