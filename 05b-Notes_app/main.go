package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	notePack "example.com/notes/notes"
	todoPack "example.com/notes/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
	saver
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo :")

	todo, err := todoPack.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(todo)

	myNote, err := notePack.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(myNote)

}

func printSomething(value interface{}) {
	// typedValue, ok := value.(int)
	// switch value.(type) {
	// case int:
	// 	fmt.Println("int:", value)
	// case float64:
	// 	fmt.Println("float:", value)
	// case string:
	// 	fmt.Println("string:", value)
	// }
}

func outputData(data displayer) {
	data.Display()

	err := saveData(data)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note has encountered an issue")
		fmt.Println(err)
		return err
	}

	fmt.Println("Your note has been saved")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Please enter the title of the note:")
	content := getUserInput("Please enter the content for the note:")

	return title, content

}

func getUserInput(dispText string) string {
	fmt.Printf("%v ", dispText)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
