package customerusecase

import (
	"github.com/BennoAlif/ps-cats-social/src/entities"
	customer "github.com/BennoAlif/ps-cats-social/src/repositories/customer"
)

type sCustomerUsecase struct {
	customerRepository customer.CustomerRepository
}

type CustomerUsecase interface {
	Create(*entities.ParamsCreateCustomer) (*entities.ResultCreateCustomer, error)
}

func New(
	customerRepository customer.CustomerRepository,
) CustomerUsecase {
	return &sCustomerUsecase{
		customerRepository: customerRepository,
	}
}
