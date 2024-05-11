package orderusecase

import (
	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sOrderUsecase) Create(details *entities.ParamsCustomerCheckout) error {
	return i.orderRepository.Create(details)
}
