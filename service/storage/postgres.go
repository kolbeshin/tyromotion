package storage

import (
	"log"
	"os"

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

	users := []models.Doctor{
		{Email: os.Getenv("USER_1"), Password: os.Getenv("PASSWORD_1")},
		{Email: os.Getenv("USER_2"), Password: os.Getenv("PASSWORD_2")},
	}

	for _, user := range users {
		result := connect.Create(&user)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
	}
	return &Storage{connect}, nil
}

func (s *Storage) GetDoctorByEmail(email string) (models.Doctor, error) {
	var doctor models.Doctor
	tx := s.db.Where("email = ?", email).First(&doctor)
	return doctor, tx.Error
}

func (s *Storage) Create(doctor models.Doctor) (models.Doctor, error) {
	tx := s.db.Create(&doctor)
	return doctor, tx.Error
}

func (s *Storage) Get(email string) (models.Doctor, error) {
	var doctor models.Doctor
	tx := s.db.First(&doctor, email)
	return doctor, tx.Error
}
