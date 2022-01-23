package context

import (
	"fmt"
	"os"
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
	indexes.AppendCreatedTodo(TodolistIndex{TodoId: todoId, Title: title, Status: "OPEN"})
	return todoId
}

func (t Todolist) ListIndexes() []*TodolistIndex {
	indexes := TodoListIndexes{t.Workdir}
	return indexes.List()
}

func (t Todolist) GetTodo(todoId int) string {
	indexOf := t.IndexOf(todoId)
	return indexOf.Show()
}

func (t Todolist) IndexOf(todoId int) *TodolistIndex {
	indexes := TodoListIndexes{t.Workdir}
	indexOf := indexes.IndexOf(todoId)
	return indexOf
}

func (t Todolist) CloseTodo(todoId int) {
	indexes := TodoListIndexes{t.Workdir}
	indexOf := indexes.IndexOf(todoId)
	indexOf.close()
	indexes.update(indexOf)
}

func (t Todolist) ModifyTodo(index *TodolistIndex) {
	indexes := TodoListIndexes{t.Workdir}
	indexOf := indexes.IndexOf(index.TodoId)
	indexOf.update(index)
	indexes.update(indexOf)
}
