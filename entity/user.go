package entity

type User struct {
	Base
	Name        string
	PhoneNumber string
	Signature   string
	Account     Account
	TimeSheets  []TimeSheet
}
