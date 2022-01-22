package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTodolistIndexes(t *testing.T) {
	indexesFile := "todoId,title,status\n-------------------\n1,todo 1,OPEN\n2,todo 2\n三,todo 3,OPEN\n"

	todolistIndexes, err := todolistIndexesParser([]byte(indexesFile))

	assert.Nil(t, err)
	assert.NotEmpty(t, todolistIndexes)
	assert.Equal(t, 1, len(todolistIndexes))
	assert.Equal(t, TodolistIndex{
		TodoId: 1,
		Title:  "todo 1",
		Status: "OPEN",
	}, *todolistIndexes[0])
}

func TestTodoListIndexes(t *testing.T) {
	todolistIndexes := []*TodolistIndex{
		{
			TodoId: 1,
			Title:  "todo 1",
			Status: "OPEN",
		}, {
			TodoId: 2,
			Title:  "todo 2",
			Status: "OPEN",
		},
	}

	todoId := latestTodoId(todolistIndexes)

	assert.Equal(t, int64(2), todoId)

	todolistIndexes = []*TodolistIndex{}

	todoId = latestTodoId(todolistIndexes)

	assert.Equal(t, int64(0), todoId)
}

func TestNewTodoId(t *testing.T) {
	testCases := []struct {
		indexesFile string
		newTodoId   int64
	}{
		{
			indexesFile: "todoId,title,status\n-------------------\n",
			newTodoId:   1,
		}, {
			indexesFile: "todoId,title,status\n-------------------\n1,todo 1,OPEN\n",
			newTodoId:   2,
		},
	}

	for _, testCase := range testCases {
		todoId := newTodoId([]byte(testCase.indexesFile))
		assert.Equal(t, testCase.newTodoId, todoId)
	}

}
