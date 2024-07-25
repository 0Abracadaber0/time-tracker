package datafixes

import (
	"time-tracker/internal/models"
)

type Trimmable interface {
	TrimRightSpaces()
}

func Trim(t *models.User) {
	t.TrimRightSpaces()
}
