package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

type questions struct {
	ques string
	ans  string
}

func main() {
	question := readQuestions("Test.csv")

	total, correct := renderQues(question)

	fmt.Println("You attemted: ", total, "\n", "Correct ans: ", correct)
}

func readQuestions(CsvFile string) []questions {

	csvFile, _ := os.Open(CsvFile)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var Testques []questions
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		}
		Testques = append(Testques, questions{
			ques: line[0],
			ans:  line[1],
		})
	}

	return Testques
}

func renderQues(Test []questions) (int, int) {
	correct, increment := 0, 0
	for i, question := range Test {
		fmt.Println("Question no: ", i)
		fmt.Println("Question: ", question.ques)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter answer: ")
		text, _ := reader.ReadString('\n')
		if strings.Trim(text, "\n") == question.ans {
			correct = correct + 1
		}
		increment = i

	}
	increment = increment + 1

	return increment, correct
}
