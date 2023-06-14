package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

type questioner interface {
	Ask()
	Check(choice string) bool
}

type singleChoiceQuestion struct {
	Question string
	Answer   string
}

func (q singleChoiceQuestion) Ask() {
	fmt.Println(q.Question)
}

func (q singleChoiceQuestion) Check(choice string) bool {
	return q.Answer == cleanUpString(choice)
}

type multipleChoiceQuestion struct {
	Question string
	Answer   string
	Options  []string
}

func (q multipleChoiceQuestion) Ask() {
	fmt.Println("Question:", q.Question)
	for i := range q.Options {
		fmt.Println(i, q.Options[i])
	}
	fmt.Println("Enter answer value")
}

func (q multipleChoiceQuestion) Check(choice string) bool {
	return q.Answer == cleanUpString(choice)
}

type quiz struct {
	Creator     string
	Items       []questioner
	TimeAllowed time.Duration
	Result      score
}

type score int

type csvResponse [][]string

type quizzer interface {
	Read(fileData csvResponse) []questioner
}

type singleChoiceQuiz struct{}

func (scq singleChoiceQuiz) Read(fileData csvResponse) []questioner {

	// parse file content return slice of [questioner]
	var examItems []questioner
	for i := range fileData {
		questionAnswerPair := fileData[i]

		item := singleChoiceQuestion{
			Question: questionAnswerPair[0],
			Answer:   cleanUpString(questionAnswerPair[1]),
		}

		examItems = append(examItems, item)
	}
	return examItems

}

type multipleChoiceQuiz struct{}

func (mcq multipleChoiceQuiz) Read(fileData csvResponse) []questioner {

	// parse file content return slice of [questioner]
	var examItems []questioner
	for i := range fileData {
		questionAnswerPair := fileData[i]

		item := multipleChoiceQuestion{
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

func createQuizFromFile(filePath, creatorName, timeAllowed string, quizType quizzer) quiz {
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

	return quiz{
		Creator:     creatorName,
		Items:       examItems,
		Result:      0,
		TimeAllowed: quizDuration,
	}

}

func (q *quiz) Run() {
	var result score
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

func (q quiz) DisplayResult() {
	fmt.Fprintf(os.Stdout, "You got %d out %d questions. \n", q.Result, len(q.Items))
}

func (q quiz) SaveToFIle(fileName string) {
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
