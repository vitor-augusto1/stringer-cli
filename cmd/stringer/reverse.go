package stringer

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vitor-augusto1/stringer-cli/pkg/stringer"
)




var reverseCmd = &cobra.Command{
  Use: "reverse",
  Aliases: []string{"rev"},
  Short: "Reverses a string",
  Args: cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    res := stringer.Reverse(args[0])
    fmt.Println(res)
  },
}


