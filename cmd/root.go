package cmd

import (
	"errors"
	"github.com/Duke1616/passkey/cmd/app"
	"github.com/Duke1616/passkey/cmd/start"
	"github.com/Duke1616/passkey/config"

	"github.com/spf13/cobra"
)

var (
	// 配置导入类型
	confType string
	// 暂未使用
	confFile string
)

var RootCmd = &cobra.Command{
	Use:   "passkey",
	Short: "Passkey通用密钥登录",
	Long:  "Passkey的原理很简单，就是在用户注册环节，可以选择生成一对密钥，分别是公钥和私钥，公钥存在服务器端，而私钥存在用户需要登录的设备上，包含但不限于电脑、手机、或者平板。",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// config 为全局变量, 只需要load 即可全局可用户
func loadGlobalConfig(configType string) error {
	// 配置加载
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
func initial() {
	err := loadGlobalConfig(confType)
	cobra.CheckErr(err)

}

func Execute(wireInit func() *app.App) {
	// 补充初始化设置
	cobra.OnInitialize(initial)

	startCmd := start.NewStartApiCommand(wireInit)
	RootCmd.AddCommand(startCmd)

	err := RootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&confType, "config-type", "t", "file", "the service config type [file/env]")
	RootCmd.PersistentFlags().StringVarP(&confFile, "config-file", "f", "config.toml", "the service config from file")
}
