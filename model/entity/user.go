package entity

type User struct {
	ID          string  `gorm:"type:varchar(255);primaryKey" json:"id"`
	Name        string  `json:"name"`
	PhoneNumber string  `json:"phone_number"`
	Account     Account `json:"account"`
}
