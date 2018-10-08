package main

import (
	"gophercises/task/db"
)

func main() {
	_, _ = db.Initialize("db/newDB")

}
