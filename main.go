package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strconv"
)

type Todo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

// Gets the text string of the todo
func (t *Todo) GetText() string {
	str := ""
	if t.Done {
		str = strconv.FormatInt(t.ID, 10) + ". [x] " + t.Name
	} else {
		str = strconv.FormatInt(t.ID, 10) + ". [ ] " + t.Name
	}
	return str
}

func updateStatus(todos *[]Todo, id int64) error {
	found := false
	for i := range *todos {
		if (*todos)[i].ID == id {
			(*todos)[i].Done = !(*todos)[i].Done
			found = true
			break
		}
	}
	if found {
		return nil
	} else {
		return fmt.Errorf("can't find %v", id)
	}
}

func main() {

	// TODO: Blob should be read from JSON file
	blob := `[{
		"id": 1,
		"name": "Todo 1",
		"done": false
	},
	{
		"id": 2,
		"name": "Todo 2",
		"done": false
	},
	{
		"id": 3,
		"name": "Todo 3",
		"done": false
	}]`

	var todos []Todo

	// Load the todos from JSON
	if err := json.Unmarshal([]byte(blob), &todos); err != nil {
		log.Fatal(err)
	}

	updateTodoFlag := flag.Int64("update", -1, "The id of todo to update")

	flag.Parse()

	// update todo status
	if *updateTodoFlag != -1 {
		if err := updateStatus(&todos, *updateTodoFlag); err != nil {
			fmt.Printf("Couldn't update todo: %v\n", err)
		}
	}

	// Always print the todos
	for _, todo := range todos {
		fmt.Println(todo.GetText())
	}
}
