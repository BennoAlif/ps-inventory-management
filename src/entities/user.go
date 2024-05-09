package entities

import (
	"time"
)

type User struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"-"`
	Token       string    `json:"token"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiredAt   time.Time `json:"-"`
}

type ParamsCreateUser struct {
	ID          int64
	Name        string
	PhoneNumber string
}
