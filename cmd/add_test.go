package cmd

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

const (
	testWorkdir = "../test_workdir"
)

func TestAddTodo(t *testing.T) {
	setUpTestWorkdir(t)

	addTodo("addTodo 1")
	addTodo("addTodo 2")

	assertTodoCreated(t, "1-addTodo 1.md")
	assertTodoCreated(t, "2-addTodo 2.md")
}

func assertTodoCreated(t *testing.T, name string) {
	files, err := ioutil.ReadDir(testWorkdir)
	assert.Nil(t, err)

	existsAddTodo1File := false
	for _, f := range files {
		if f.Name() == name {
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
	err := initTodolist()
	if err != nil {
		assert.Nil(t, err)
	}
}
