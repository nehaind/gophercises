package cmd

import (
	"fmt"
	"gophercises/task/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list down whole task list.",
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := db.ListTask()
		if err != nil {
			fmt.Println("error found.")
		}

		if len(tasks) < 1 {
			fmt.Println("no element in the list")

		}
		for _, task := range tasks {
			fmt.Println("index: ", task.Key, " value: ", task.Value)
		}

	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
