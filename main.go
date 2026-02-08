package main

import (
	"errors"
	"fmt"
)

func main() {
	getUserInput("note title")
	getUserInput("note content")
}

func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	var value string
	fmt.Scan(&value)
	if value == "" {
		// fmt.Println("Please enter a value")
		return "", errors.New("you should put some value here ")
	}
	return value, nil
}
