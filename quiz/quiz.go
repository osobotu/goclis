package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type QuestionItem struct {
	Question string
	Answer   string
}
type Quiz struct {
	Creator string
	Items   []QuestionItem
	Result  Score
}

type Score int

func createQuizFromFile(filePath, creatorName string) Quiz {
	// open csv file
	file, err := os.Open(filePath)
	check(err)

	// read file content
	fileData, err := csv.NewReader(file).ReadAll()
	check(err)

	// parse file content to create a quiz
	var examItems []QuestionItem
	for i := range fileData {
		questionAnswerPair := fileData[i]

		item := QuestionItem{
			Question: questionAnswerPair[0],
			Answer:   cleanUpString(questionAnswerPair[1]),
		}

		examItems = append(examItems, item)
	}

	return Quiz{
		Creator: creatorName,
		Items:   examItems,
		Result:  0,
	}

}

func (q *Quiz) Run() {
	var result Score
	fmt.Println("Exam created by: ", q.Creator)

	for _, item := range q.Items {
		fmt.Println(item.Question)

		// get user answer
		var givenAnswer string
		fmt.Scanln(&givenAnswer)

		// compare with correct answer
		givenAnswer = cleanUpString(givenAnswer)
		if givenAnswer == item.Answer {
			result += 1
		}

	}
	q.Result = result

}

func (q Quiz) DisplayResult() {
	fmt.Fprintf(os.Stdout, "You got %d out %d questions. \n", q.Result, len(q.Items))
}

func (q Quiz) SaveToFIle(fileName string) {
	message := fmt.Sprintf("Created by: %v\nYou got %d out of %d questions", q.Creator, q.Result, len(q.Items))
	os.WriteFile(fileName, []byte(message), 0666)
}

// helpers
func check(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func cleanUpString(s string) string {
	s = strings.ToLower(s)
	s = strings.Trim(s, " ")
	return s
}
