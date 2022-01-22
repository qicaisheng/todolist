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
	TodoId int64
	title  string
	status string
}

type TodoListIndexes struct {
	Workdir string
}

func latestTodoId(indexes []*TodolistIndex) int64 {
	if len(indexes) == 0 {
		return 0
	}
	return indexes[len(indexes)-1].TodoId
}

func todolistIndexesParser(bytes []byte) ([]*TodolistIndex, error) {
	todolistIndexesText := strings.TrimSpace(string(bytes))
	lines := strings.Split(todolistIndexesText, "\n")
	var todolistIndexes []*TodolistIndex
	for i, line := range lines[2:] {
		split := strings.Split(line, ",")
		if len(split) != 3 {
			log.Printf(".todolist_index file is broken, the content \"%s\" is ignored", line)
			continue
		}
		parseInt, err := strconv.ParseInt(split[0], 0, 0)
		if err != nil {
			log.Printf(".todolist_index file is broken, the content \"%s\" is ignored", line)
			continue
		}
		todolistIndexes = append(todolistIndexes, &TodolistIndex{
			TodoId: parseInt,
			title:  split[1],
			status: split[2],
		})
		log.Printf("%v", *todolistIndexes[i])
	}
	return todolistIndexes, nil
}

func (util TodoListIndexes) NewTodoId() int64 {
	indexesFile := util.indexesFile()
	file, err := os.ReadFile(indexesFile)
	if err != nil {
		fmt.Println("read indexes file error: ", err)
		os.Exit(1)
	}
	return newTodoId(file)
}

func newTodoId(indexesFile []byte) int64 {
	indexes, err := todolistIndexesParser(indexesFile)
	if err != nil {
		os.Exit(1)
	}
	return latestTodoId(indexes) + 1
}

func (util TodoListIndexes) InitTodolistIndexesFile() error {
	indexesFile := util.indexesFile()
	if _, err := os.Stat(indexesFile); os.IsNotExist(err) {
		err := os.WriteFile(indexesFile, []byte(todolistIndexesFileTemplate), 0644)
		return err
	}
	return nil
}

func (util TodoListIndexes) indexesFile() string {
	return filepath.Join(util.Workdir, todolistIndexesFileName)
}
