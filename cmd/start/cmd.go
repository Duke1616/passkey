package start

import (
	"github.com/Duke1616/passkey/cmd/app"
	"github.com/spf13/cobra"
)

func NewStartApiCommand(wireInit func() *app.App) *cobra.Command {
	Cmd := &cobra.Command{
		Use:   "start",
		Short: "passkey API服务",
		Long:  "passkey 启动",
		RunE: func(cmd *cobra.Command, args []string) error {
			server := wireInit()
			err := server.Web.Run(":8100")
			if err != nil {
				return err
			}
			return nil
		},
		SilenceUsage: true,
	}

	return Cmd
}
