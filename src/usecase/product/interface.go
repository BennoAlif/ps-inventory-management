package productusecase

import (
	"github.com/BennoAlif/ps-cats-social/src/entities"
	product "github.com/BennoAlif/ps-cats-social/src/repositories/product"
)

type sProductUsecase struct {
	productRepository product.ProductRepository
}

type ProductUsecase interface {
	Create(*entities.ParamsCreateProduct) (*ResultCreate, error)
	FindMany(*entities.ProductSearchFilter) ([]*entities.Product, error)
	Update(*string, *entities.ParamsUpdateProduct) error
	Delete(*string) error
}

func New(productRepository product.ProductRepository) ProductUsecase {
	return &sProductUsecase{
		productRepository: productRepository,
	}
}
