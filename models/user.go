package models

type User struct {
	UserID       uint64 `json:"user_id,omitempty"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	Email        string `json:"email,omitempty"`
	Gender       int    `json:"gender,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
