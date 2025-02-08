package models

import "time"

type User struct {
	ID               int       `json:"id"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	Role             string    `json:"role"`
	RegistrationDate time.Time `json:"registration_date"`
}

func NewUser(username, email, password, role string) *User {
	return &User{
		Username:         username,
		Email:            email,
		Password:         password,
		Role:             role,
		RegistrationDate: time.Now(),
	}
}
