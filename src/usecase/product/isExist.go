package productusecase

import "github.com/BennoAlif/ps-cats-social/src/entities"

func (i *sProductUsecase) IsExist(productId *string) error {
	filters := entities.ProductSearchFilter{
		ID: *productId,
	}

	product, _ := i.productRepository.IsExists(&filters)

	if !product {
		return ErrProductNotFound
	}

	return nil
}
