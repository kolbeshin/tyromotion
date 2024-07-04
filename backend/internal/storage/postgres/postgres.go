package postgres

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"tyromotion/backend/internal/models"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(dsn string) (*Storage, error) {
	connect, err := gorm.Open(postgres.Open(dsn), nil)
	if err != nil {
		return nil, err
	}

	err = connect.AutoMigrate(&models.Patient{})
	if err != nil {
		return nil, err
	}
	err = connect.AutoMigrate(&models.PatientInfo{})
	if err != nil {
		return nil, err
	}
	err = connect.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	password1, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("PASSWORD_1")), 14)
	password2, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("PASSWORD_2")), 14)

	users := []models.User{
		{Email: os.Getenv("USER_1"), Password: password1, PhoneNumber: os.Getenv("PHONE_NUMBER_1")},
		{Email: os.Getenv("USER_2"), Password: password2, PhoneNumber: os.Getenv("PHONE_NUMBER_2")},
	}

	for _, user := range users {
		result := connect.Create(&user)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
	}
	return &Storage{connect}, nil
}

func (s *Storage) RegisterUser(user models.User) (models.User, error) {
	tx := s.db.Create(&user)
	return user, tx.Error
}

func (s *Storage) LoginUser(user models.User) (models.User, error) {
	tx := s.db.Where("email = ?", user.Email).First(&user)
	return user, tx.Error
}

func (s *Storage) GetUser(claims *jwt.StandardClaims) (models.User, error) {
	var user models.User
	tx := s.db.Where("id = ?", claims.Issuer).First(&user)
	return user, tx.Error
}

func (s *Storage) GetAllPatientsFromTable() ([]models.Patient, error) {
	var patient []models.Patient
	result := s.db.Find(&patient)
	if result.Error != nil {
		return nil, result.Error
	}
	return patient, nil
}

func (s *Storage) CreatePatient(patient models.Patient) (models.Patient, error) {
	tx := s.db.Create(&patient)
	return patient, tx.Error
}

func (s *Storage) GetCompletedTreatments(id int) (models.Patient, error) {
	var patient models.Patient
	tx := s.db.First(&patient, id)
	return patient, tx.Error
}
