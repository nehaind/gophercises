package cmd

import (
	"fmt"
	"gophercises/task/db"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var docmd = &cobra.Command{
	Use:   "do",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		indexString := strings.Join(args, " ")
		index, err := strconv.Atoi(indexString)
		if err != nil {

			fmt.Printf("error occured")
			return
		}
		errOccurred := db.DeleteTask(index)
		if errOccurred == nil {
			fmt.Println("marked the task as done and removed from the queue")

		}
	},
}

func init() {
	RootCmd.AddCommand(docmd)
}
