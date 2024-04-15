package models

type Doctor struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Patient struct {
	ID              uint64 `json:"id"`
	FullName        string `json:"fullname"`
	DateOfBirthday  string `json:"birthday"`
	Sex             string `json:"sex"`
	DateOfDischarge string `json:"dateofdischarge"`
	InsuranceNumber string `json:"policy"`
	Archive         string `json:"archive"`
	Info            string `json:"info"`
}
