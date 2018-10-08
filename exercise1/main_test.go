package main

import (
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	main()
	dashtest.ControlCoverage(m)
}

func TestRenderQues(t *testing.T) {
	var ques []questions

	ques = append(ques, questions{
		ques: "one",
		ans:  "one",
	})
	ques = append(ques, questions{
		ques: "two",
		ans:  "2",
	})
	renderQues(ques)
}

// func TestReadQuestions(t *testing.T) {
// 	inputs := []string{"Test.csv", "test1.csv"}
// 	for _, input := range inputs {
// 		_, error := readQuestions(input)
// 		if error != nil {
// 			t.Error("error occured while test case : ", error)
// 		}
// 	}
// }
