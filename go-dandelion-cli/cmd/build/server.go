package build

import (
	"fmt"
	"github.com/gly-hub/go-dandelion/go-dandelion-cli/internal/build"
	"github.com/spf13/cobra"
)

var (
	appName  string
	StartCmd = &cobra.Command{
		Use:          "build",
		Short:        "生成服务结构代码",
		Example:      "go-dandelion-cli build -n example-application",
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
	StartCmd.PersistentFlags().StringVarP(&appName, "name", "n", "example-server", "服务器名称")
}

func setup() {
}

func run() error {
	fmt.Print("Type of service you want to create, enter a number（1-rpc 2-http）:")
	var serverType int
	if _, err := fmt.Scanln(&serverType); err != nil {
		fmt.Println("An error occurred while reading the input:", err)
		return nil
	}
	switch serverType {
	case 1:
		return build.Rpc(appName)
	case 2:
		return build.Http(appName)
	}

	return nil
}
