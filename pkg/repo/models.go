package repo

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	message string
}
