package main

import (
	"bytes"
	"os"
	"testing"
)

const testFile = "./testdata/words_test.txt"

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4

	res := count(b, false, false)
	if res != exp {
		t.Errorf("Expected %d, but got %d instead.\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2\n word3 word4 word5\n word6 word7")
	exp := 3

	res := count(b, true, false)
	if res != exp {
		t.Errorf("Expected %d, but got %d instead.\n", exp, res)
	}
}

func TestCountBytes(t *testing.T) {
	s := "word1 word2\n word3 word4 word5\n word6 word7"
	b := bytes.NewBufferString(s)
	exp := len(string(s))

	res := count(b, false, true)
	if res != exp {
		t.Errorf("Expected %d, but got %d instead.\n", exp, res)
	}
}

func TestCountLinesFromFile(t *testing.T) {

	file, err := os.Open(testFile)
	if err != nil {
		t.Errorf("Could not open test file")
	}
	exp := 4

	res := count(file, true, false)
	if res != exp {
		t.Errorf("Expected %d, but got  %d instead.\n", exp, res)
	}

}
func TestCountBytesFromFile(t *testing.T) {

	file, err := os.Open(testFile)
	if err != nil {
		t.Errorf("Could not open test file")
	}
	s := "hello\nworld\nhello\nworld"
	exp := len(s)

	res := count(file, false, true)
	if res != exp {
		t.Errorf("Expected %d, but got  %d instead.\n", exp, res)
	}

}
