package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	filePath = "./problems.csv"
)

func main() {

	// ! flags
	quizFilename := flag.String("file", "", "Quiz file name")
	// save := flag.String("save", "", "save saves quiz result to this file")
	flag.Parse()

	// ! get user inputted filename
	fmt.Println(*quizFilename)
	if *quizFilename != "" {
		filePath = "./" + *quizFilename
		// filePath = filepath.Join(".", filePath)
	}
	fmt.Println(filePath)

	// err := run(filePath, *save)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	names := []string{"Suraj", "Steve"}
	fmt.Println("Original:", names)
	for i := range names {
		fmt.Println(i)
		names[i] = strings.ToUpper(names[i])
		fmt.Println(names[i])
	}
	fmt.Println("New:", names)
}

// TODO: modify the quiz to take multiple files
// TODO: modify quiz to log results to output file
// TODO: modify quiz to include a timer that defaults to 30secs but can be override.

func check(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func run(filePath, outFilename string) error {
	// ! open a csv file
	openFile, err := os.Open(filePath)
	check(err)

	// ! read file content
	fileData, err := csv.NewReader(openFile).ReadAll()
	check(err)

	// ! parse file content
	numberOfQuestions := len(fileData)
	numberOfCorrectAnswers := 0

	for i := range fileData {
		questionAnswerPair := fileData[i]
		question := questionAnswerPair[0]

		// ! clean up answer from csv file
		answer := strings.ToLower(questionAnswerPair[1])
		answer = strings.Trim(answer, " ")
		fmt.Println(question)

		// ! get the user input
		var givenAnswer string
		fmt.Scanln(&givenAnswer)

		// ! compare with correct answer
		givenAnswer = strings.ToLower(givenAnswer)
		if givenAnswer == answer {
			numberOfCorrectAnswers++
		}
	}

	if outFilename != "" {
		// ! save to specified file
		fmt.Println("save to file")
		// // create or open file
		// f, err := os.OpenFile(*&outFilename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		// if err != nil {
		// 	log.Fatal(err)
		// 	os.Exit(1)
		// }
		// defer f.Close()

	} else {
		// ! print user's on stdout
		message := fmt.Sprintf("You got %d out of %d questions", numberOfCorrectAnswers, numberOfQuestions)
		fmt.Println(message)
	}
	return nil
}
