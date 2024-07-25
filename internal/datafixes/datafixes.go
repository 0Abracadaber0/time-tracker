package datafixes

import (
	"strings"
	"time-tracker/internal/models"
	"unicode"
)

func TrimRightSpace(user *models.User) {
	user.Id = strings.TrimRightFunc(user.Id, unicode.IsSpace)
	user.Surname = strings.TrimRightFunc(user.Surname, unicode.IsSpace)
	user.Name = strings.TrimRightFunc(user.Name, unicode.IsSpace)
	user.Patronymic = strings.TrimRightFunc(user.Patronymic, unicode.IsSpace)
	user.Address = strings.TrimRightFunc(user.Address, unicode.IsSpace)
}
