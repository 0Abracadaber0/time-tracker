package models

import (
	"strings"
	"unicode"
)

type User struct {
	Id         string `json:"id"`
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}

func (user *User) TrimRightSpaces() {
	user.Id = strings.TrimRightFunc(user.Id, unicode.IsSpace)
	user.Surname = strings.TrimRightFunc(user.Surname, unicode.IsSpace)
	user.Name = strings.TrimRightFunc(user.Name, unicode.IsSpace)
	user.Patronymic = strings.TrimRightFunc(user.Patronymic, unicode.IsSpace)
	user.Address = strings.TrimRightFunc(user.Address, unicode.IsSpace)
}
