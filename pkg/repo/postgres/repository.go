package postgres

import (
	"crud_example/pkg/api"
	models "crud_example/pkg/repo"
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgRepository struct {
	db *gorm.DB
}

func NewPgRepository(DbDSN string) *PgRepository {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: DbDSN, PreferSimpleProtocol: true}), &gorm.Config{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&models.Log{})
	if err != nil {
		return nil
	}
	return &PgRepository{db: db}
}

func (r *PgRepository) Add(schema *api.LogPostSchema) error {
	log := models.Log{Message: schema.Message}
	err := r.db.Create(&log)
	if err != nil {
		return errors.New("WTF")
	}
	return nil
}

func (r *PgRepository) Get(id int) (*api.LogGetSchema, error) {
	log := models.Log{}
	if err := r.db.Where("id = ?", id).First(&log); errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("not found")
	}
	return &api.LogGetSchema{ID: id, Message: log.Message, DtCreated: log.CreatedAt}, nil
}

func (r *PgRepository) Delete(id int) error {
	r.db.Delete(&models.Log{}, id)
	return nil
}
