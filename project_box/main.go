package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// This is a simple project to practice the concepts of Go
// I learnt from the crash course. From basic data types to
// concurrency.

// I will be writing a simple cli app with the following features.
// 1. Take user input - first name, last name, matriculation number,
// 	department and details about their project.
// 	Project details should include
//  topic name, supervisor and boolean flag of whether it is group
// 	or individual.
// 2. Store the details in a csv file.

var (
	schoolName = "Federal University of Technology, Minna"
	session    = "2021/2022"
)

type Student struct {
	firstName           string
	lastName            string
	matriculationNumber string
}

type Project struct {
	topic          string
	supervisorName string
}

func main() {
	run()
}

func run() {
	greetUser()
	printDivider()

	var student Student
	// var project Project

	firstName, lastName, matriculationNumber := getUserPersonalDetails()
	student.firstName = firstName
	student.lastName = lastName
	student.matriculationNumber = matriculationNumber

	printDivider()

	greetUserAccordingToTimeOfDay(firstName)
	// fmt.Println("Enter your project details below")

	// printDivider()

	// topic, supervisorName := getUserProjectDetails()
	// project.topic = topic
	// project.supervisorName = supervisorName

}

func greetUser() {
	fmt.Printf("Welcome to %v's Project Box for %v.\n", schoolName, session)
	fmt.Println("Enter your personal details to begin.")
}

func printDivider() {
	fmt.Println("============================================")
}

func greetUserAccordingToTimeOfDay(name string) {
	now := time.Now()
	var greetingPrefix string
	switch {
	case 0 > now.Hour() && now.Hour() < 12:
		greetingPrefix = "Good morning"
	case 12 <= now.Hour() && now.Hour() <= 17:
		greetingPrefix = "Good afternoon"
	default:
		greetingPrefix = "Good evening"
	}
	fmt.Println(greetingPrefix, name)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

	}

}

func getUserPersonalDetails() (string, string, string) {
	var firstName string
	var lastName string
	var matriculationNumber string

	fmt.Println("Enter your first name")
	fmt.Scanln(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scanln(&lastName)
	fmt.Println("Enter your matriculation number")
	fmt.Scanln(&matriculationNumber)

	return firstName, lastName, matriculationNumber
}

func getUserProjectDetails() (string, string) {
	var topic string
	var supervisorName string
	fmt.Println("Enter your project topic")
	fmt.Scan(&topic)
	fmt.Println("Enter your supervisor's name")
	fmt.Scan(&supervisorName)
	return topic, supervisorName
}
