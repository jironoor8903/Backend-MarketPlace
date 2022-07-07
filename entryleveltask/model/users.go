package model

type User struct {
	ID          int64       `json:"id"`
	Username     string    `json:"username"`
	Email     string    `json:"email"`
	PasswordHash    string    `json:"passwordhash"`
}
