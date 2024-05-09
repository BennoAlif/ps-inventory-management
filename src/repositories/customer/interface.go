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
	// FindMany(*entities.ParamsFindManyCustomer) ([]*entities.Customer, error)
	IsExists(*entities.FiltersCustomer) (bool, error)
}

func New(db *sql.DB) CustomerRepository {
	return &sCustomerRepository{DB: db}
}
