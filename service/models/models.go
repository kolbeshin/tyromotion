package models

type Doctor struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Patient struct {
	FullName        string `json:"fullname"`
	ID              uint64 `json:"id"`
	DateOfBirthday  string `json:"birthday"`
	InsuranceNumber string `json:"policy"`
	Archive         string `json:"archive"`
}
