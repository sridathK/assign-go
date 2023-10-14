package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrNothingInsideFile = errors.New("nothing is inside file")

func main() {
	data, err := os.ReadFile("sample.txt")
	if err == nil {
		fmt.Println(CountWords(string(data)))
	} else {
		fmt.Println(err)
	}
}

func CountWords(data string) (int, error) {
	if len(data) == 0 {
		return 0, ErrNothingInsideFile
	}
	chars := []rune(data)
	var count int
	for i := 0; i < len(chars); i++ {
		if chars[i] == ' ' {
			count++
		}
	}
	return count, nil
}
