package productusecase

import "github.com/BennoAlif/ps-cats-social/src/entities"

func (i *sProductUsecase) TotalPrice(details []entities.ProductDetails) (int, error) {
	return i.productRepository.TotalPrice(details)
}
