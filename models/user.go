package models

type User struct {
	UserID       uint
	UserName     string
	Password     string
	Email        string
	Gender       int
	AccessToken  string
	RefreshToken string
}
