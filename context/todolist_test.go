package context

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
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
	todoId2 := todoList.AddTodo("addTodo 2")

	assertTodoCreated(t, todoId1, "addTodo 1")
	assertTodoCreated(t, todoId2, "addTodo 2")

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

	todo1 := todoList.GetTodo(1)
	assert.Equal(t, "# 1-addTodo 1\n## status\nOPEN\n", todo1)

	teardownTestWorkdir(t)
}

func teardownTestWorkdir(t *testing.T) {
	err := os.RemoveAll(testWorkdir)
	assert.Nil(t, err)
}

func assertTodoCreated(t *testing.T, todoId int, name string) {
	files, err := ioutil.ReadDir(testWorkdir)
	assert.Nil(t, err)

	fileName := fmt.Sprintf("%v-%s.md", todoId, name)
	existsAddTodo1File := false
	for _, f := range files {
		if f.Name() == fileName {
			existsAddTodo1File = true
			break
		}
	}

	assert.True(t, existsAddTodo1File)
}

func setUpTestWorkdir(t *testing.T) {
	viper.Set("workdir", testWorkdir)
	if _, err := os.Stat(testWorkdir); os.IsNotExist(err) {
		err := os.Mkdir(testWorkdir, os.ModePerm)
		assert.Nil(t, err)
	}
}
