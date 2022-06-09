package scripts

import (
	"github.com/armnerd/go-skeleton/cmd/scripts/app/demo"
	"github.com/spf13/cobra"
)

var (
	Register = &cobra.Command{
		Use: "cmd",
	}
)

func init() {
	Register.AddCommand(demo.Register)
}
