package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:  "done",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("ID inválido: %s", args[0])
		}
		return s.Done(id)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
