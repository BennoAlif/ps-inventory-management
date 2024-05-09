package entities

import (
	"time"
)

type Customer struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}

type ParamsCreateCustomer struct {
	Name        string `json:"name" validate:"required,min=5,max=50"`
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16"`
}

type ResultCreateCustomer struct {
	ID          int64  `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
}

type FiltersCustomer struct {
	ID          int64
	Name        string
	PhoneNumber string
}
