package context

import (
	"strconv"
)

type TodolistIndex struct {
	TodoId int
	Title  string
	Status string
}

func (i TodolistIndex) String() string {
	return strconv.Itoa(i.TodoId) + "," + i.Title + "," + i.Status
}

func (i *TodolistIndex) close() {
	i.Status = "CLOSED"
}

func (i *TodolistIndex) update(index *TodolistIndex) {
	i.Title = index.Title
	i.Status = index.Status
}
