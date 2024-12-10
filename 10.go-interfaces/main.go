package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
	Display()
}

// type displayer interface{
// 	Display()
// }
// Embeded interfaces
// type ouputtable interface{
// 	saver
// 	displayer
// }

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

	err = outputData(todo, "todo")
	if err != nil {
		return
	}

	err = outputData(userNote, "note")
	if err != nil {
		return
	}
	printAnything(1)
	printAnything(1.5)
	printAnything("Vikas")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

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

func outputData(data saver, text string) error {
	data.Display()
	return saveData(data, text)
}

func saveData(data saver, text string) error {
	err := data.Save()

	if err != nil {
		fmt.Printf("Saving the %v failed.\n", text)
		return err
	}

	fmt.Printf("Saving the %v succeeded!\n", text)
	return nil
}

func printAnything(value interface{}) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer Value: ", intVal)
		return
	}

	strVal, ok := value.(string)
	if ok {
		fmt.Println("String Value: ", strVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float Value: ", floatVal)
		return
	}

	// //switch case for type of values
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer Value: ", value)
	// case float64:
	// 	fmt.Println("Float Value: ", value)
	// case string:
	// 	fmt.Print("String value: ", value)
	// default:
	// 	fmt.Print(value)
	// }
}
