package utils

import (
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

type TodolistIndexes struct {
	Indexes []*TodolistIndex
}

func (i TodolistIndexes) LatestTodoId() int64 {
	indexes := i.Indexes
	return indexes[len(indexes)-1].TodoId
}

func todolistIndexesParser(bytes []byte) (TodolistIndexes, error) {
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
	return TodolistIndexes{todolistIndexes}, nil
}

func InitTodolistIndexesFile(workdir string) error {
	indexesFile := filepath.Join(workdir, todolistIndexesFileName)
	if _, err := os.Stat(indexesFile); os.IsNotExist(err) {
		err := os.WriteFile(indexesFile, []byte(todolistIndexesFileTemplate), 0644)
		return err
	}
	return nil
}
