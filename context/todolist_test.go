package context

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	testWorkdir = "../test_workdir"
)

func TestAll(t *testing.T) {
	setUpTestWorkdir(t)
	todoList := Todolist{Workdir: testWorkdir}

	todoList.InitTodolist()

	todolistIndexes := todoList.ListIndexes()
	assert.Empty(t, todolistIndexes)

	todoId1 := todoList.AddTodo("addTodo 1")
	assert.Equal(t, 1, todoId1)
	todoId2 := todoList.AddTodo("addTodo 2")
	assert.Equal(t, 2, todoId2)

	todolistIndexes = todoList.ListIndexes()
	assert.NotEmpty(t, todolistIndexes)
	assert.Equal(t, 2, len(todolistIndexes))
	assert.Equal(t, TodolistIndex{
		TodoId: 1,
		Title:  "addTodo 1",
		Status: "OPEN",
	}, *todolistIndexes[0])
	assert.Equal(t, TodolistIndex{
		TodoId: 2,
		Title:  "addTodo 2",
		Status: "OPEN",
	}, *todolistIndexes[1])

	todoList.CloseTodo(1)
	todo1 := todoList.GetTodo(1)
	assert.Equal(t, "# 1-addTodo 1\n## status\nCLOSED\n", todo1)

	todoList.ModifyTodo(&TodolistIndex{
		TodoId: 2,
		Title:  "updated todo 2",
		Status: "OPEN",
	})
	todo2 := todoList.GetTodo(2)
	assert.Equal(t, "# 2-updated todo 2\n## status\nOPEN\n", todo2)

	teardownTestWorkdir(t)
}

func teardownTestWorkdir(t *testing.T) {
	err := os.RemoveAll(testWorkdir)
	assert.Nil(t, err)
}

func setUpTestWorkdir(t *testing.T) {
	viper.Set("workdir", testWorkdir)
	if _, err := os.Stat(testWorkdir); os.IsNotExist(err) {
		err := os.Mkdir(testWorkdir, os.ModePerm)
		assert.Nil(t, err)
	}
}
