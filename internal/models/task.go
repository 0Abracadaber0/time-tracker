package models

import (
	"strings"
	"time"
	"unicode"
)

type Task struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Time      int       `json:"time"`
	LastStart time.Time `json:"last_start"`
	IsWorking bool      `json:"is_working"`
}

func (task *Task) TrimRightSpaces() {
	task.Id = strings.TrimRightFunc(task.Id, unicode.IsSpace)
	task.Name = strings.TrimRightFunc(task.Name, unicode.IsSpace)
}
