package main

import (
	"bufio"
	"fmt"
	"noteApp/note"
	"noteApp/todo"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo.DisplayNote()
	err = todo.Save()

	if err != nil {
		fmt.Println("saving the todo is faild ")
		return
	}

	fmt.Println("todo save succeeded!")

	userNote.DisplayNote()
	err = userNote.Save()

	if err != nil {
		fmt.Println("saving the note is faild ")
		return
	}

	fmt.Println("note save succeeded!")

}

// this func is totaly a shit and learn a lot
func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("note title:")
	content := getUserInput("note content:")
	return title, content
}
