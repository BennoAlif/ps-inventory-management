package entities

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	SKU         string    `json:"sku"`
	Category    string    `json:"category"`
	ImageUrl    string    `json:"imageUrl"`
	Stock       int       `json:"stock"`
	Notes       string    `json:"notes"`
	Price       int       `json:"price"`
	Location    string    `json:"location"`
	IsAvailable bool      `json:"isAvailable"`
	CreatedAt   time.Time `json:"createdAt"`
}

type CreateProduct struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

type ProductSearchFilter struct {
	ID          string      // Limit the output based on the product id
	Limit       int         // Limit the output of the data
	Offset      int         // Offset for pagination
	Name        string      // Filter based on product's name
	IsAvailable interface{} // Filter if the product is available
	Category    string      // Filter based on category
	SKU         string      // Filter based on product's SKU
	Price       string      // Sort by price
	InStock     interface{} // Check whether the stock is > 0
	CreatedAt   string      // Sort by created time
}

type ParamsCreateProduct struct {
	ID          uuid.UUID
	Name        string
	Sku         string
	Category    string
	Description string
	ImageUrl    string
	Notes       string
	Price       int
	Stock       int
	Location    string
	IsAvailable bool
}

type ParamsUpdateProduct struct {
	Name        string // Product's name (not null, minLength 1, maxLength 30)
	Sku         string // Product's SKU (not null, minLength 1, maxLength 30)
	Category    string // Product's category (not null, enum of: "Clothing", "Accessories", "Footwear", "Beverages")
	Notes       string // Product's notes (not null, minLength 1, maxLength 200)
	ImageUrl    string // Product's image URL (not null, should be URL)
	Price       int    // Product's price (not null, min: 1)
	Stock       int    // Product's stock (not null, min: 1, max: 100000)
	Location    string // Product's location (not null, minLength 1, maxLength 200)
	IsAvailable bool   // Product's availability (not null)
}
