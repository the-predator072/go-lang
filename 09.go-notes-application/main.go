package main

import (
	"fmt"

	"example.com/note/note"
)

func main() {
	userNote, err := note.New()
	if err != nil {
		panic(err)
	}
	userNote.Display()
	err = userNote.Save()
	if err != nil {
		panic("saving the note failed")
	}
	fmt.Println("Note saved successfully.")
}
