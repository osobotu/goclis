package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	fileName := flag.String("file", "", "Filename to count")
	flag.Parse()
	// Printing out the number of words entered in the stdin using count

	if *fileName != "" {
		files := []string{*fileName}

		files = append(files, flag.Args()...)
		fmt.Println(files)

		// count the words or lines in each file
		for i := range files {
			file, err := os.Open(files[i])
			if err != nil {
				os.Exit(1)
			}
			fmt.Println(count(file, *lines, *bytes))
		}

	} else {
		fmt.Println(count(os.Stdin, *lines, *bytes))
	}

}

// func countMultipleFiles(files []string) int {

// }

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}
	return wc
}
