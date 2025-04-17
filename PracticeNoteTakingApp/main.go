package main

import (
	"bufio"
	"examples.com/NotetakingApp/note"
	"examples.com/NotetakingApp/todo"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

//type displayer interface {
//	Display()
//}

//type outputtable interface {
//	Save() error
//	Display()
//}

type outputtable interface {
	saver
	Display()
	//DoSomething(int) string
}

func main() {
	//It is dangerous though
	printSomething("Some string")
	printSomething(1)
	printSomething(1.5)

	result := addGenerics(1, 2)
	fmt.Println(result)

	title, content := getNoteData()
	todoText := getTodoData()
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Output the userNote
	err = outputData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Saved Note successfully")

	////Display Todo
	//todo.Display()
	////Save the todo
	//err = saveData(todo)
	//if err != nil {
	//	return
	//}

	//Output the todo
	err = outputData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Saved Todo successfully")

}

func getTodoData() string {
	return getUserInput("Todo text: ")
}

// value can be of any type
// Even Println() accepts any type
func printSomething(value interface{}) {
	intVal, ok := value.(int)
	if ok {
		intVal = intVal + 1
		fmt.Println("Integer: ", intVal)
		return
	}
	floatVal, ok := value.(float64)
	if ok {
		floatVal = floatVal + 1
		fmt.Println("Float: ", intVal)
		return
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String: ", stringVal)
		return
	}

	//switch value.(type) {
	//case int:
	//	fmt.Println("Integer: ", value)
	//case string:
	//	fmt.Println("String: ", value)
	//case float64:
	//	fmt.Println("Float: ", value)
	//default:
	//	fmt.Println("Unknown type: ", value)
	//}
}
func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
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

func add(a, b interface{}) any {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}

	aStr, aIsStr := a.(string)
	bStr, bIsStr := b.(string)

	if aIsStr && bIsStr {
		return aStr + bStr
	}

	return a
}

func addGenerics[T int | float64 | string](a, b T) T {
	return a + b
}
