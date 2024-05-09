package customerusecase

import (
	"github.com/BennoAlif/ps-cats-social/src/entities"
	"github.com/BennoAlif/ps-cats-social/src/helpers"
	userusecase "github.com/BennoAlif/ps-cats-social/src/usecase/user"
)

func (i *sCustomerUsecase) Create(p *entities.ParamsCreateCustomer) (*entities.ResultCreateCustomer, error) {
	emailMx := helpers.IsValidPhoneNumber(p.PhoneNumber)

	if emailMx != nil {
		return nil, emailMx
	}

	filters := entities.FiltersCustomer{
		PhoneNumber: p.PhoneNumber,
	}

	checkPhoneNumber, _ := i.customerRepository.IsExists(&filters)

	if checkPhoneNumber {
		return nil, userusecase.ErrPhoneNumberAlreadyUsed
	}

	customer, err := i.customerRepository.Create(p)
	if err != nil {
		return nil, err
	}

	result := &entities.ResultCreateCustomer{
		ID:          customer.ID,
		Name:        customer.Name,
		PhoneNumber: customer.PhoneNumber,
	}

	return result, nil
}
