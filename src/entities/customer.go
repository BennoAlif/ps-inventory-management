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

type ProductDetails struct {
	ProductId string `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
}

type ParamsCustomerCheckout struct {
	CustomerId     string         `json:"customerId" validate:"required"`
	ProductDetails ProductDetails `json:"productDetails"`
	Paid           int            `json:"paid" validate:"required,min=1"`
	Change         int            `json:"change" validate:"required,min=0"`
}

type ResultCreateCustomer struct {
	ID          int64  `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
}

type ParamsCustomer struct {
	ID          string
	Name        string
	PhoneNumber string
}
