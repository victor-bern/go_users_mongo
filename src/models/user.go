package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
}
