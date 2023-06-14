package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

type Questioner interface {
	Ask()
	Check(choice string) bool
}

type SingleChoiceQuestion struct {
	Question string
	Answer   string
}

func (q SingleChoiceQuestion) Ask() {
	fmt.Println(q.Question)
}

func (q SingleChoiceQuestion) Check(choice string) bool {
	return q.Answer == cleanUpString(choice)
}

type MultipleChoiceQuestion struct {
	Question string
	Answer   string
	Options  []string
}

func (q MultipleChoiceQuestion) Ask() {
	fmt.Println("Question:", q.Question)
	for i := range q.Options {
		fmt.Println(i, q.Options[i])
	}
	fmt.Println("Enter answer value")
}

func (q MultipleChoiceQuestion) Check(choice string) bool {
	return q.Answer == cleanUpString(choice)
}

type Quiz struct {
	Creator     string
	Items       []Questioner
	TimeAllowed time.Duration
	Result      Score
}

type Score int

type Quizzer interface {
	Read(fileData [][]string) []Questioner
}

type SingleChoiceQuiz struct {
	Creator     string
	Items       []Questioner
	TimeAllowed time.Duration
	Result      Score
}

func (scq SingleChoiceQuiz) Read(fileData [][]string) []Questioner {

	// parse file content return slice of [Questioner]
	var examItems []Questioner
	for i := range fileData {
		questionAnswerPair := fileData[i]

		item := SingleChoiceQuestion{
			Question: questionAnswerPair[0],
			Answer:   cleanUpString(questionAnswerPair[1]),
		}

		examItems = append(examItems, item)
	}
	return examItems

}

type MultipleChoiceQuiz struct{}

func (mcq MultipleChoiceQuiz) Read(fileData [][]string) []Questioner {

	// parse file content return slice of [Questioner]
	var examItems []Questioner
	for i := range fileData {
		questionAnswerPair := fileData[i]

		item := MultipleChoiceQuestion{
			Question: questionAnswerPair[0],
			Answer:   cleanUpString(questionAnswerPair[1]),
			Options: []string{
				// TODO: randomize the options
				cleanUpString(questionAnswerPair[2]),
				cleanUpString(questionAnswerPair[1]),
				cleanUpString(questionAnswerPair[3]),
				cleanUpString(questionAnswerPair[4]),
			},
		}

		examItems = append(examItems, item)
	}
	return examItems

}

func createQuizFromFile(filePath, creatorName, timeAllowed string, quizType Quizzer) Quiz {
	// open csv file
	file, err := os.Open(filePath)
	check(err)

	// close file
	defer file.Close()

	// read file content
	fileData, err := csv.NewReader(file).ReadAll()
	check(err)

	// read fileData and get questions
	examItems := quizType.Read(fileData)

	// assign allocated time or set to default (30secs)
	quizDuration := time.Duration(30)
	if timeAllowed != "" {
		quizDuration, err = time.ParseDuration(timeAllowed)
		check(err)
	}

	return Quiz{
		Creator:     creatorName,
		Items:       examItems,
		Result:      0,
		TimeAllowed: quizDuration,
	}

}

func (q *Quiz) Run() {
	var result Score
	fmt.Println("Exam created by: ", q.Creator)

	for _, item := range q.Items {
		// fmt.Println(item.Question)
		item.Ask()

		// get user answer
		var givenAnswer string
		fmt.Scanln(&givenAnswer)

		// compare with correct answer
		givenAnswer = cleanUpString(givenAnswer)
		if item.Check(givenAnswer) {
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
		fmt.Println("Error: ", err)
		return err
	}
	return nil
}

func cleanUpString(s string) string {
	s = strings.ToLower(s)
	s = strings.Trim(s, " ")
	return s
}
