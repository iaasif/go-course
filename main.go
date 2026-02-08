package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNoteData()
	if err != nil {
		println(err)
		return
	}
	fmt.Println(title, content)
}

func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	var value string
	fmt.Scan(&value)
	if value == "" {
		return "", errors.New("Invalid input.")
	}
	return value, nil
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("note title")
	if err != nil {
		println(err)
		return "", "", err
	}
	content, err := getUserInput("note content")
	if err != nil {
		println(err)
		return "", "", err
	}
	return title, content, nil
}
