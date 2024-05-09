package customerrepository

import (
	"database/sql"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

type sCustomerRepository struct {
	DB *sql.DB
}

type CustomerRepository interface {
	Create(*entities.ParamsCreateCustomer) (*entities.Customer, error)
	FindMany(*entities.ParamsCustomer) ([]*entities.Customer, error)
	IsExists(*entities.ParamsCustomer) (bool, error)
}

func New(db *sql.DB) CustomerRepository {
	return &sCustomerRepository{DB: db}
}
