package main

import (
	"os"

	"github.com/Vitorio-Pereira/task_cli/internal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
