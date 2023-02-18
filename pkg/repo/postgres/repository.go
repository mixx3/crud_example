package postgres

import (
	"gorm.io/gorm"
	"log"
)

type pgRepository struct {
	logger log.Logger
	db     *gorm.DB
}
