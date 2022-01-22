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
	setUp(t)

	addTodo("addTodo 1")

	files, err := ioutil.ReadDir(testWorkdir)
	assert.Nil(t, err)

	existsAddTodo1File := false
	for _, f := range files {
		if f.Name() == "addTodo 1.md" {
			existsAddTodo1File = true
			break
		}
	}

	assert.True(t, existsAddTodo1File)

}

func setUp(t *testing.T) {
	viper.Set("workdir", testWorkdir)
	if _, err := os.Stat(testWorkdir); os.IsNotExist(err) {
		err := os.Mkdir(testWorkdir, os.ModePerm)
		assert.Nil(t, err)
	}
}
