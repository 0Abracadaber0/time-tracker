package models

type User struct {
	Id         string `json:"id"`
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}
