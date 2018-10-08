package db

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type task struct {
	Key   int
	Value string
}

//Initialize will init the DB
func Initialize(dbPath string) (*bolt.DB, error) {
	var err error

	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 2 * time.Second})
	if err != nil {
		return nil, err
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte(taskBucket))
		return err
	})
	return db, nil
}

//AddTask is used to input Value in db
func AddTask(task string) (int, error) {
	var id int
	if task == "" || task == " " {
		fmt.Println("invalid task name")
		return -1, nil
	}
	err := db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = (int(id64))
		return b.Put([]byte(Itob(id)), []byte(task))

	})

	return id, err
}

//ListTask returns the list of all task
func ListTask() ([]task, error) {
	var task1 []task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			task1 = append(task1, task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	return task1, err
}

//DeleteTask is meant to delte item using the Value
func DeleteTask(Key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(Itob(Key))
	})
}

//Itob function
func Itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	fmt.Println(b)
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
