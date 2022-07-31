package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Level    int    `json:"level"`
	Token    string `json:"token"`
}
