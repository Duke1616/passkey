package main

import (
	"errors"
	"passkey-demo/config"
)

func main() {
	err := loadGlobalConfig("file")
	if err != nil {
		panic("加载配置文件出错")
	}

	server := InitWebServer()

	server.web.Run(":8100")
}

func loadGlobalConfig(configType string) error {
	switch configType {
	case "file":
		err := config.LoadConfigFromToml()
		if err != nil {
			return err
		}
	case "env":
		err := config.LoadConfigFromEnv()
		if err != nil {
			return err
		}
	default:
		return errors.New("unknown config type")
	}

	return nil
}
