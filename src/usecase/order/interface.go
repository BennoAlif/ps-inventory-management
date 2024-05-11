package orderusecase

import (
	"github.com/BennoAlif/ps-cats-social/src/entities"

	orderepository "github.com/BennoAlif/ps-cats-social/src/repositories/order"
)

type sOrderUsecase struct {
	orderRepository orderepository.OrderRepository
}

type OrderUsecase interface {
	Create(*entities.ParamsCustomerCheckout) error
	// FindMany(*entities.OrderSearchFilter) ([]*entities.Order, error)
}

func New(orderRepository orderepository.OrderRepository) OrderUsecase {
	return &sOrderUsecase{
		orderRepository: orderRepository,
	}
}
