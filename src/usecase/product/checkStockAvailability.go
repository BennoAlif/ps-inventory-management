package productusecase

import "github.com/BennoAlif/ps-cats-social/src/entities"

func (i *sProductUsecase) CheckStockAvailability(details []*entities.ProductDetails) (bool, error) {
	return i.productRepository.CheckStockAvailability(details)
}
