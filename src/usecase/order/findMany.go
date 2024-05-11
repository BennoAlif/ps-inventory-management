package orderusecase

import (
	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sOrderUsecase) FindMany(filters *entities.SearchOrderFilter) ([]*entities.ParamsCustomerCheckout, error) {
	return i.orderRepository.FindMany(filters)
}
