package todos

import (
	"encoding/json"
	"log"
	"os"
)

type Item struct {
	Text  string
	Done  bool
	Index int
}

func AddItems(args []string, dataFile string) []Item {
	todos := GetItems(dataFile)
	for _, x := range args {
		todos = append(todos, Item{Text: x})
	}
	SaveItems(dataFile, todos)
	return todos
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
		return []Item{}
	}
	var items []Item
	err = json.Unmarshal(marshal, &items)
	if err != nil {
		log.Fatal(err)
	}
	return items
}

// func CheckItems(args []string, dataFile string) {
// 	todos := GetItems(dataFile)
// 	if len(args) == 0 {
// 		todos[0].Done = true
// 	} else {
// 		for _, todo := range todos {
// 			for _, index := range
// 			if i == todo.Index {
// 				todo.Done = true
// 			}
// 		}
// 	}
// 	fmt.Println(todos)
// }

func ClearItems(filename string) {
	err := os.WriteFile(filename, []byte("[]"), 0644)
	if err != nil {
		log.Fatal("Unable to Clear list")
	}
}
