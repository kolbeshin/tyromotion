package storage

import (
	"github.com/SenyashaGo/tyromotion/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(dsn string) (*Storage, error) {
	connect, err := gorm.Open(postgres.Open(dsn), nil)
	if err != nil {
		return nil, err
	}

	err = connect.AutoMigrate(&models.Doctor{})
	if err != nil {
		return nil, err
	}
	err = connect.AutoMigrate(&models.Patient{})
	if err != nil {
		return nil, err
	}
	return &Storage{connect}, nil
}

func (s *Storage) Create(patient models.Patient) (int64, error) {
	tx := s.db.Create(&patient)
	return tx.RowsAffected, tx.Error
}

func (s *Storage) Get(id uint64) (models.Doctor, error) {
	var doctor models.Doctor
	tx := s.db.First(&doctor, id)
	return doctor, tx.Error
}
