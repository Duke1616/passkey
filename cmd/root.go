package cmd

import (
	"github.com/Duke1616/passkey/cmd/app"
	"github.com/Duke1616/passkey/cmd/start"
	"github.com/Duke1616/passkey/config"

	"github.com/spf13/cobra"
)

var (
	// 暂时未使用
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
func loadGlobalConfig() error {
	// 优先使用环境变量，环境变量格式 PASSKEY_MYSQL_USER=root
	// 如果环境不存在，读取配置文件
	err := config.LoadConfig()
	if err != nil {
		return err
	}

	return nil
}

func initial() {
	err := loadGlobalConfig()
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
	RootCmd.PersistentFlags().StringVarP(&confFile, "config-file", "f", "config.toml", "the service config from file")
}
