package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title string, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content cannot be empty")
	}
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
func (n *Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ReplaceAll(fileName, "/", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(n)
	
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
func (n *Note) Display() {
	fmt.Printf("Title: %s\nContent: %s\nCreated At: %s", n.Title, n.Content, n.CreatedAt.Format(time.RFC1123))
}

