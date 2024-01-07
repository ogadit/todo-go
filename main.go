package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

// - Read Todos from file

type Todo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func (t Todo) print() string {
	str := ""
	if t.Done {
		str = strconv.FormatInt(t.ID, 10) + ". [x] " + t.Name
	} else {
		str = strconv.FormatInt(t.ID, 10) + ". [ ] " + t.Name
	}
	return str
}

func main() {
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

	if err := json.Unmarshal([]byte(blob), &todos); err != nil {
		log.Fatal(err)
	}

	for _, todo := range todos {
		fmt.Println(todo.print())
	}
}
