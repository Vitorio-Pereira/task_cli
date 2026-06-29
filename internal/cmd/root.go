package cmd

import (
	"github.com/Vitorio-Pereira/task_cli/internal/store"
	"github.com/spf13/cobra"
)

var s *store.Store

var rootCmd = &cobra.Command{
	Use:   "task_cli",
	Short: "This is a To Do for Cli, very easy and fast to use on the day",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		s = store.New()
		return s.Load()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
