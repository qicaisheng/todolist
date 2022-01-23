package context

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strconv"
)

type Todolist struct {
	Workdir string
}

func (t Todolist) InitTodolist() {
	indexes := TodoListIndexes{t.Workdir}
	err := indexes.InitTodolistIndexesFile()
	if err != nil {
		fmt.Printf("init toolist error: %v\n", err)
		os.Exit(1)
	}
}

func (t Todolist) AddTodo(title string) int {
	indexes := TodoListIndexes{t.Workdir}
	todoId := indexes.NewTodoId()
	fileName := fmt.Sprintf("%v-%s.md", strconv.Itoa(todoId), title)
	filePath := filepath.Join(viper.GetString("workdir"), fileName)
	todoItem := fmt.Sprintf("# %v-%s\n## status\n%s\n", strconv.Itoa(todoId), title, "OPEN")
	err := os.WriteFile(filePath, []byte(todoItem), 0644)
	if err != nil {
		_ = fmt.Errorf("write file error: %v", err)
	}
	indexes.AppendCreatedTodo(TodolistIndex{TodoId: todoId, Title: title, Status: "OPEN"})
	return todoId
}

func (t Todolist) ListIndexes() []*TodolistIndex {
	indexes := TodoListIndexes{t.Workdir}
	return indexes.List()
}

func (t Todolist) GetTodo(todoId int) string {
	indexes := TodoListIndexes{t.Workdir}
	indexOf := indexes.IndexOf(todoId)
	fileName := fmt.Sprintf("%v-%s.md", todoId, indexOf.Title)

	todoItemFile := filepath.Join(t.Workdir, fileName)
	todoDetail, err := os.ReadFile(todoItemFile)
	if err != nil {
		fmt.Printf("get todo detail error: %v\n", err)
		os.Exit(1)
	}
	return string(todoDetail)
}
