package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTodolistIndexes(t *testing.T) {
	indexesFile := "todoId,title,status\n-------------------\n1,todo 1,OPEN\n2,todo 2\n三,todo 3,OPEN\n"

	todolistIndexes, err := todolistIndexesParser([]byte(indexesFile))

	assert.Nil(t, err)
	assert.NotNil(t, todolistIndexes)
	assert.NotEmpty(t, todolistIndexes.Indexes)
	assert.Equal(t, 1, len(todolistIndexes.Indexes))
	assert.Equal(t, TodolistIndex{
		TodoId: 1,
		title:  "todo 1",
		status: "OPEN",
	}, *todolistIndexes.Indexes[0])
}

func TestTodoListIndexes(t *testing.T) {
	todolistIndexes := TodolistIndexes{Indexes: []*TodolistIndex{
		{
			TodoId: 1,
			title:  "todo 1",
			status: "OPEN",
		}, {
			TodoId: 2,
			title:  "todo 2",
			status: "OPEN",
		},
	}}

	latestTodoId := todolistIndexes.LatestTodoId()

	assert.Equal(t, int64(2), latestTodoId)

	todolistIndexes = TodolistIndexes{Indexes: []*TodolistIndex{}}

	latestTodoId = todolistIndexes.LatestTodoId()

	assert.Equal(t, int64(0), latestTodoId)
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
