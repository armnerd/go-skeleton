package demo

import (
	"github.com/armnerd/go-skeleton/cmd/scripts/common"
	"github.com/spf13/cobra"
)

var (
	Register = &cobra.Command{
		Use: "demo",
	}
)

var (
	hello = &cobra.Command{
		Use: "hello",
		PreRun: func(cmd *cobra.Command, args []string) {
			common.Depend()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return helloHandler(args)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			common.Release()
		},
	}
	world = &cobra.Command{
		Use: "world",
		RunE: func(cmd *cobra.Command, args []string) error {
			return worldHandler(args)
		},
	}
)

func init() {
	// 注册子命令
	Register.AddCommand(hello)
	Register.AddCommand(world)
}
