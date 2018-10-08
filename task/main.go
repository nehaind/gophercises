package main

import (
	"fmt"
	"gophercises/task/cmd"
	"gophercises/task/db"
)

func main() {
	_, err := db.Initialize("db/newDB")
	cmd.RootCmd.Execute()
	if err != nil {
		fmt.Println("error found")
	}

}
