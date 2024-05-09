package customerusecase

import "github.com/BennoAlif/ps-cats-social/src/entities"

func (i *sCustomerUsecase) FindMany(filters *entities.ParamsCustomer) ([]*entities.Customer, error) {
	customers, err := i.customerRepository.FindMany(filters)

	if err != nil {
		return nil, err
	}

	return customers, nil
}
