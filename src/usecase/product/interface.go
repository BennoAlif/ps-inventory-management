package productusecase

import (
	"github.com/BennoAlif/ps-cats-social/src/entities"
	customerrepository "github.com/BennoAlif/ps-cats-social/src/repositories/customer"
	orderepository "github.com/BennoAlif/ps-cats-social/src/repositories/order"
	productrepository "github.com/BennoAlif/ps-cats-social/src/repositories/product"
)

type sProductUsecase struct {
	productRepository   productrepository.ProductRepository
	orderRepository     orderepository.OrderRepository
	cusrtomerRepository customerrepository.CustomerRepository
}

type ProductUsecase interface {
	Create(*entities.ParamsCreateProduct) (*ResultCreate, error)
	FindMany(*entities.ProductSearchFilter) ([]*entities.Product, error)
	IsExist(*string) error
	IsExistMany([]*string) error
	Update(*string, *entities.ParamsUpdateProduct) error
	Delete(*string) error
	TotalPrice([]entities.ProductDetails) (int, error)
	Checkout(*entities.ParamsCustomerCheckout) error
	CheckStockAvailability([]*entities.ProductDetails) (bool, error)
}

func New(
	productRepository productrepository.ProductRepository,
	orderRepository orderepository.OrderRepository,
	cusrtomerRepository customerrepository.CustomerRepository,
) ProductUsecase {
	return &sProductUsecase{
		productRepository:   productRepository,
		orderRepository:     orderRepository,
		cusrtomerRepository: cusrtomerRepository,
	}
}
