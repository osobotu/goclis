package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"flag"
	"todo"
)

var todoFileName = ".todo.json"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed by Obotu\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2023\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage information:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "You can add new tasks as a command line argument or via then STDIN stream.\n")
		flag.PrintDefaults()
	}

	// ! parsing command line flags
	add := flag.Bool("add", false, "Add task to the todo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	delete := flag.Int("delete", 0, "Item to be deleted")
	verbose := flag.Bool("v", false, "enable verbose output")
	uncompleted := flag.Bool("uc", false, "show only uncompleted ToDos")

	flag.Parse()

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		showError(err)
	}

	switch {
	case *list:
		if *verbose {
			output := ""
			for i, v := range *l {
				prefix := "   "
				if v.Done {
					prefix = "X  "
				}
				formattedCreatedAt := v.CreatedAt.Format("02 January 2006 03:04:05 PM")

				output += fmt.Sprintf("%s%d: %s createdAt: %s\n", prefix, i+1, v.Task, formattedCreatedAt)
			}
			fmt.Print(output)
			return
		} else if *uncompleted {
			output := "Uncompleted ToDos:\n"
			for i, v := range *l {
				prefix := "   "
				if v.Done {
					prefix = "X  "
					continue
				}
				formattedCreatedAt := v.CreatedAt.Format("02 January 2006 03:04:05 PM")

				output += fmt.Sprintf("%s%d: %s createdAt: %s\n", prefix, i+1, v.Task, formattedCreatedAt)

			}
			fmt.Print(output)
			return
		}
		// ! List current Todo items
		fmt.Print(l)

	case *complete > 0:
		// ! Complete the given time
		if err := l.Complete(*complete); err != nil {
			showError(err)
		}
		// ! save the new list
		if err := l.Save(todoFileName); err != nil {
			showError(err)
		}
	case *add:
		task, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			showError(err)
		}
		l.Add(task)

		if err := l.Save(todoFileName); err != nil {
			showError(err)
		}
	case *delete > 0:
		if err := l.Delete(*delete); err != nil {
			showError(err)
		}

		if err := l.Save(todoFileName); err != nil {
			showError(err)
		}

	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be blank")
	}
	return s.Text(), nil
}

func showError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
