package note

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// struct with tags for adding meta data
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func New() (Note, error) {
	title, content, err := getNoteData()
	if err != nil {
		return Note{}, err
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) Display() {
	fmt.Printf("You note titled %v\n has content\n %v \n", n.Title, n.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	os.WriteFile(fileName, json, 0644)
	return nil
}

func getNoteData() (string, string, error) {
	title, err1 := getUserInput("Note Title: ")
	if err1 != nil {
		return "", "", err1
	}
	content, err2 := getUserInput("Note Content: ")
	if err2 != nil {
		return "", "", err2
	}
	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text, nil
}
