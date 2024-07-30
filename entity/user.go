package entity

type User struct {
	Base
	Name        string  `json:"name"`
	PhoneNumber string  `json:"phone_number"`
	Signature   string  `json:"signature_url"`
	Account     Account `json:"account"`
	TimeSheets  []TimeSheet
}
