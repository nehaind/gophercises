package cmd

import (
	"github.com/spf13/cobra"
)

//RootCmd is used to initiate the task
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
}
