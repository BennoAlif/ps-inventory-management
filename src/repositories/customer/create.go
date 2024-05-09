package customerrepository

import (
	"log"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sCustomerRepository) Create(p *entities.ParamsCreateCustomer) (*entities.Customer, error) {
	var id int64
	err := i.DB.QueryRow("INSERT INTO customers (name, phone_number) VALUES ($1, $2) RETURNING id", p.Name, p.PhoneNumber).Scan(&id)
	if err != nil {
		log.Printf("Error inserting customer: %s", err)
		return nil, err
	}

	customer := &entities.Customer{
		ID:          id,
		Name:        p.Name,
		PhoneNumber: p.PhoneNumber,
	}

	return customer, nil
}
