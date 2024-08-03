package repository

type AccountRepository interface {
	AccountActivation(email string) error
}
