package notes

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() {
	title, content := getNoteData()
	note, err := New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}
	note.Display()
	err = note.Save()
	if err != nil {
		fmt.Println("Error saving note to file:", err)
		return
	}
}

func getNoteData() (string, string){
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getUserInput(prompt string) (string) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")
	return strings.TrimSpace(value)
}