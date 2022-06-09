package demo

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Register = &cobra.Command{
		Use: "demo",
	}
)

var (
	hello = &cobra.Command{
		Use:    "hello",
		PreRun: func(cmd *cobra.Command, args []string) {},
		RunE: func(cmd *cobra.Command, args []string) error {
			return helloHandler()
		},
	}
	world = &cobra.Command{
		Use:    "world",
		PreRun: func(cmd *cobra.Command, args []string) {},
		RunE: func(cmd *cobra.Command, args []string) error {
			return worldHandler()
		},
	}
)

func init() {
	// 注册子命令
	Register.AddCommand(hello)
	Register.AddCommand(world)
}

func helloHandler() error {
	fmt.Println("Hello")
	return nil
}

func worldHandler() error {
	fmt.Println("world")
	return nil
}
