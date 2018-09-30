package cmd

import (
	"fmt"
	"gophercises/task/db"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		index, _ := db.AddTask(task)
		if index == -1 {
			fmt.Println("Wrong input")
			return
		}
		fmt.Println("value added at index: ", index)

	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
