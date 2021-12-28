package todos

import (
	"encoding/json"
	"log"
	"os"
)

type Item struct {
	Text string
}

func SaveItems(filename string, items []Item) {
	marshalled, err := json.Marshal(items)
	if err != nil {
		log.Fatal("Failed to JSON encode todo list")
	}
	err = os.WriteFile(filename, marshalled, 0644)
	if err != nil {
		log.Fatal("Failed to Write to ", filename)
	}
}

func GetItems(filename string) []Item {
	marshal, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var items []Item
	err = json.Unmarshal(marshal, &items)
	if err != nil {
		log.Fatal(err)
	}
	return items
}

func GetLocation() string {
	dataFile, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not locate User Home Directory")
	}
	return dataFile + string(os.PathSeparator) + ".todos.json"
}
