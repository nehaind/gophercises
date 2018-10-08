package cmd

import (
	"fmt"
	"gophercises/task/db"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	dbc, _ := db.Initialize("C:\\Users\\gs-1553\\go\\src\\gophercises\\task\\db\\newDB")
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{"Complete Exercise"}
	addCmd.Run(addCmd, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "value added at index:")
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	fmt.Println(string(content))
	file.Close()
	dbc.Close()
}

func TestAddInvalid(t *testing.T) {
	dbc, _ := db.Initialize("C:\\Users\\gs-1553\\go\\src\\gophercises\\task\\db\\newDB")
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{""}
	addCmd.Run(addCmd, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "Wrong input")
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	fmt.Println(string(content))
	file.Close()
	dbc.Close()
}
