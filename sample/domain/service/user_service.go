package service

// UserService user domain service
type UserService interface {
	Duplicated(username string) error
}
