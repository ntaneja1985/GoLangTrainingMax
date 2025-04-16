package main

import (
	"bufio"
	"examples.com/NotetakingApp/note"
	"fmt"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Display User Note
	userNote.Display()

	//Save the file
	err = userNote.Save()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Saved Note successfully")

}

func getNoteData() (string, string) {
	title := getUserInput("Enter the Note Title")
	//if err != nil {
	//	fmt.Println(err)
	//	return "", "", err
	//}
	content := getUserInput("Enter the Note Content")
	//if err != nil {
	//	fmt.Println(err)
	//	return "", "", err
	//}

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	//Create a reader that listens on the command line
	reader := bufio.NewReader(os.Stdin)
	//rune in Go is a single character
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	//fmt.Scan(&value)
	//if value == "" {
	//	return "", errors.New("User input is empty")
	//}
	return text
}
