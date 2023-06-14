package main

import (
	"flag"
	"path/filepath"
	"strings"
)

var (
	filePath = "./problems.csv"
)

func main() {

	// ! flags
	quizFilename := flag.String("file", "", "Quiz file name to read questions from")
	quizType := flag.String("type", "single", "Pass quiz type. 'single' for single choice. 'multiple' for multiple choice")
	save := flag.String("save", "", "Saves quiz result to given file")

	flag.Parse()

	//  get user inputted filename
	if *quizFilename != "" {
		filePath = filepath.Join(".", *quizFilename)
	}

	var exam quiz
	if strings.ToLower(*quizType) == "multiple" {

		exam = createQuizFromFile(filePath, "Stephen", "30s", multipleChoiceQuiz{})
	} else {

		exam = createQuizFromFile(filePath, "Stephen", "30s", singleChoiceQuiz{})
	}

	exam.Run()

	if *save != "" {
		exam.SaveToFIle(*save)
	} else {
		exam.DisplayResult()
	}

}

// TODO: modify the quiz to take multiple files
// TODO: modify quiz to include a timer that defaults to 30secs but can be override.
