package todos

import (
	"os"
	"testing"
	"time"
)

func TestClearItems(t *testing.T) {
	currentTime := time.Now()
	filename := "/tmp/tst-todo-clear-" + currentTime.Format("01-02-2006")
	os.WriteFile(filename, []byte("123"), 0644)
	ClearItems(filename)
	contents, err := os.ReadFile(filename)
	if err != nil {
		t.Errorf("Unable to read file %s, after clear", filename)
	}

	if string(contents) != "[]" {
		t.Errorf("Contents are not [] but %s", contents)
	}
}

func TestAddItems(t *testing.T) {
	currentTime := time.Now()
	filename := "/tmp/tst-todo-add-" + currentTime.Format("01-02-2006")
	ClearItems(filename)
	todos := AddItems(
		[]string{"Test", "Ing"},
		filename)
	if len(todos) < 2 {
		t.Error("Todos were not added to the pile")
	} else if len(todos) > 2 {
		t.Error("Todos were more than expected")
	}
}

func TestSavedTodos(t *testing.T) {
	currentTime := time.Now()
	filename := "/tmp/tst-todo-add-" + currentTime.Format("01-02-2006")
	ClearItems(filename)
	todos := AddItems(
		[]string{"Test", "Ing"},
		filename)
	savedTodos := GetItems(filename)

	if len(todos) != len(savedTodos) {
		t.Error("Todos were not saved correctly")
		t.Errorf("Expected: %#v\nReceived: %#v", todos, savedTodos)
	}
}

func TestDoubleItems(t *testing.T) {
	currentTime := time.Now()
	filename := "/tmp/tst-todo" + currentTime.Format("01-02-2006")
	ClearItems(filename)
	todos := AddItems(
		[]string{"Test", "Ing"},
		filename)
	todos = AddItems(
		[]string{"Test", "Ing", "Ly"},
		filename)
	if len(todos) < 5 {
		t.Errorf("Adding two times gives less then expected\n\n %#v", todos)
	} else if len(todos) > 5 {
		t.Errorf("Adding two times gives were more than expected\n\n %#v", todos)
	}

}
