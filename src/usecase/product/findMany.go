package productusecase

import (
	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sProductUsecase) FindMany(filters *entities.ProductSearchFilter) ([]*entities.Product, error) {
	allProducts, err := i.productRepository.FindMany(filters)

	if err != nil {
		return nil, err
	}

	return allProducts, nil
}
