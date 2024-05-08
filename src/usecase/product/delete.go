package productusecase

import "github.com/BennoAlif/ps-cats-social/src/entities"

func (i *sProductUsecase) Delete(productId *string) error {
	filters := entities.ProductSearchFilter{
		ID: *productId,
	}

	product, _ := i.productRepository.IsExists(&filters)

	if !product {
		return ErrProductNotFound
	}

	err := i.productRepository.Delete(productId)

	if err != nil {
		return err
	}

	return nil
}
