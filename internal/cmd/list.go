package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listenCmd = &cobra.Command{
	Use: "list",
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks := s.List()
		for _, task := range tasks {
			if task.Done {
				fmt.Printf("%d [X] %s\n", task.Id, task.Title)
			} else {
				fmt.Printf("%d [ ] %s\n", task.Id, task.Title)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listenCmd)
}
