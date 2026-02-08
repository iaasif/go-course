package main

import (
	"fmt"
	"noteApp/note"
)

func main() {
	title, content := getNoteData()
	userNote, error := note.New(title, content)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(userNote)
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)
	var value string
	fmt.Scan(&value)

	return value
}

func getNoteData() (string, string) {
	title := getUserInput("note title")
	content := getUserInput("note content")
	return title, content
}
