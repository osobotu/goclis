package main

import (
	"flag"
	"path/filepath"
)

var (
	filePath = "./problems.csv"
)

func main() {

	// ! flags
	quizFilename := flag.String("file", "", "Quiz file name to read questions from")
	save := flag.String("save", "", "Saves quiz result to given file")

	flag.Parse()

	//  get user inputted filename
	if *quizFilename != "" {
		filePath = filepath.Join(".", filePath)
	}

	exam := createQuizFromFile(filePath, "Stephen")
	exam.Run()

	if *save != "" {
		exam.SaveToFIle(*save)
	} else {
		exam.DisplayResult()
	}

}

// TODO: modify the quiz to take multiple files
// TODO: modify quiz to include a timer that defaults to 30secs but can be override.
