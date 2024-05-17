package models

import "gorm.io/gorm"

type Doctor struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model
	Name        string `json:"name"`
	PhoneNumber string `json:"phone number" gorm:"unique"`
	Email       string `json:"email" gorm:"unique"`
	Password    []byte `json:"-"`
}

type Patient struct {
	gorm.Model
	FullName        string `json:"fullname"`
	DateOfBirthday  string `json:"birthday"`
	Sex             string `json:"sex"`
	DateOfDischarge string `json:"dateOfDischarge"`
	InsuranceNumber string `json:"policy"`
	Archive         string `json:"archive"`
	Info            string `json:"info"`
}

type PatientInfo struct {
	gorm.Model
	DateOfTreatment string `json:"dateOfTreatment"`
	Time            string `json:"time"`
	Duration        uint16 `json:"duration"`
	Type            string `json:"type"`
	Device          string `json:"device"`
	Comment         string `json:"comment"`
}
