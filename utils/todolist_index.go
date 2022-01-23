package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	todolistIndexesFileName     = ".todolist_indexes"
	todolistIndexesFileTemplate = "todoId,title,status\n-------------------\n"
)

type TodolistIndex struct {
	TodoId int
	Title  string
	Status string
}

func (i TodolistIndex) String() string {
	return strconv.Itoa(i.TodoId) + "," + i.Title + "," + i.Status + "\n"
}

type TodoListIndexes struct {
	Workdir string
}

func latestTodoId(indexes []*TodolistIndex) int {
	if len(indexes) == 0 {
		return 0
	}
	return indexes[len(indexes)-1].TodoId
}

func todolistIndexesParser(bytes []byte) ([]*TodolistIndex, error) {
	todolistIndexesText := strings.TrimSpace(string(bytes))
	lines := strings.Split(todolistIndexesText, "\n")
	var todolistIndexes []*TodolistIndex
	for _, line := range lines[2:] {
		split := strings.Split(line, ",")
		if len(split) != 3 {
			fmt.Printf(".todolist_index file is broken, the content \"%s\" is ignored", line)
			continue
		}
		parseInt, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Printf(".todolist_index file is broken, the content \"%s\" is ignored", line)
			continue
		}
		todolistIndexes = append(todolistIndexes, &TodolistIndex{
			TodoId: parseInt,
			Title:  split[1],
			Status: split[2],
		})
	}
	return todolistIndexes, nil
}

func (indexes TodoListIndexes) NewTodoId() int {
	indexesFile := indexes.indexesFile()
	file, err := os.ReadFile(indexesFile)
	if err != nil {
		fmt.Println("read indexes file error: ", err)
		os.Exit(1)
	}
	return newTodoId(file)
}

func newTodoId(indexesFile []byte) int {
	indexes, err := todolistIndexesParser(indexesFile)
	if err != nil {
		os.Exit(1)
	}
	return latestTodoId(indexes) + 1
}

func indexMap(indexes []*TodolistIndex) map[int]*TodolistIndex {
	todolistIndexMap := map[int]*TodolistIndex{}
	for _, index := range indexes {
		todolistIndexMap[index.TodoId] = index
	}
	return todolistIndexMap
}

func (indexes TodoListIndexes) InitTodolistIndexesFile() error {
	indexesFile := indexes.indexesFile()
	if _, err := os.Stat(indexesFile); os.IsNotExist(err) {
		err := os.WriteFile(indexesFile, []byte(todolistIndexesFileTemplate), 0644)
		return err
	}
	return nil
}

func (indexes TodoListIndexes) AppendCreatedTodo(index TodolistIndex) {
	indexesFile := indexes.indexesFile()
	f, err := os.OpenFile(indexesFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)
	if _, err := f.WriteString(index.String()); err != nil {
		log.Println(err)
	}
}

func (indexes TodoListIndexes) IndexOf(todoId int) *TodolistIndex {
	list := indexes.List()
	return indexMap(list)[todoId]
}

func (indexes TodoListIndexes) List() []*TodolistIndex {
	indexesFile := indexes.indexesFile()
	file, err := os.ReadFile(indexesFile)
	if err != nil {
		fmt.Println("read indexes file error: ", err)
		os.Exit(1)
	}
	todolistIndexes, err := todolistIndexesParser(file)
	if err != nil {
		os.Exit(1)
	}

	return todolistIndexes
}

func (indexes TodoListIndexes) indexesFile() string {
	return filepath.Join(indexes.Workdir, todolistIndexesFileName)
}
