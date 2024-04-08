package models

type Doctor struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Patient struct {
	FullName        string `json:"fullname"`
	DateOfBirthday  string `json:"birthday"`
	InsuranceNumber string `json:"policy"`
	Archive         string `json:"archive"`
}
