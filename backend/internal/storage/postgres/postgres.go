package postgres

import (
	"github.com/golang-jwt/jwt"
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

	err = connect.AutoMigrate(&models.Doctor{})
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

func (s *Storage) GetDoctorByEmail(email string) (models.Doctor, error) {
	var doctor models.Doctor
	tx := s.db.Where("email = ?", email).First(&doctor)
	return doctor, tx.Error
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

func (s *Storage) CreateDoctor(doctor models.Doctor) (models.Doctor, error) {
	tx := s.db.Create(&doctor)
	return doctor, tx.Error
}

func (s *Storage) Get(email string) (models.Doctor, error) {
	var doctor models.Doctor
	tx := s.db.First(&doctor, email)
	return doctor, tx.Error
}
