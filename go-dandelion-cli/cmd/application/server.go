package application

import (
	"github.com/gly-hub/go-dandelion-cli/internal/build"
	"github.com/spf13/cobra"
)

var (
	appName  string
	StartCmd = &cobra.Command{
		Use:          "app",
		Short:        "生成服务结构代码",
		Example:      "go-dandelion-cli app -n example-application",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&appName, "name", "n", "example-server", "应用名称")
}

func setup() {
}

func run() error {
	build.BuildApplication(appName)
	return nil
}
