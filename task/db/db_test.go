package db

import (
	"fmt"
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
func TestInitializeInvalidDbPath(t *testing.T) {
	_, error := Initialize("C:\\Users\\gs-1554\\my.db")

	if error != nil {
		t.Error("error found in the function", error)
	}
	//db.Close()
}

func TestInitialize(t *testing.T) {
	_, error := Initialize("my.db")
	if error != nil {
		t.Error("error found in the function", error)
	}

}

func TestAddTask(t *testing.T) {

	taskList := [4]string{
		"add",
		"add new",
	}
	var err error
	for _, task := range taskList {
		_, err = AddTask(task)
	}

	//fmt.Println("the i is used", i)
	if err != nil {
		t.Errorf("Task is not added: %d", err)
	}
}
func TestItob(t *testing.T) {
	variable := 12678687
	by := Itob(variable)
	if by == nil {
		t.Error("invalid next sequence value")
	}
	fmt.Printf("by: %T\n", variable)
	fmt.Printf("by: %T\n", by)

}

func TestBtoi(t *testing.T) {
	byte := []byte{0, 0, 0, 0, 0, 0, 0, 10}
	integerValue := btoi(byte)
	if integerValue == 0 {
		t.Error("invalid next sequence value")
	}

}

func TestListTask(t *testing.T) {

	task, err := ListTask()
	if err != nil {

		t.Error("error found", err)
	}
	fmt.Println("value of task", task)
	fmt.Println("value of error", err)
}

func TestDeleteTask(t *testing.T) {

	taskList := []int{
		1,
		2,
		3,
		4,
	}
	for _, task := range taskList {
		err := DeleteTask(task)
		if err != nil {

			t.Error("error found", err)
		}
	}

}
