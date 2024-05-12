package productusecase

import (
	"strconv"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sProductUsecase) Checkout(p *entities.ParamsCustomerCheckout) error {
	customerExist, _ := i.cusrtomerRepository.IsExists(&entities.ParamsCustomer{
		ID: p.CustomerID,
	})

	if !customerExist {
		return ErrCustomerNotFound
	}

	var productIDs []*string

	for _, detail := range p.ProductDetails {
		_, err := strconv.Atoi(detail.ProductID)
		if err != nil {
			return ErrProductNotFound
		}
		productIDs = append(productIDs, &detail.ProductID)
	}

	productDetails := make([]*entities.ProductDetails, len(p.ProductDetails))
	for i, detail := range p.ProductDetails {
		productDetails[i] = &detail
	}

	err := i.IsExistMany(productIDs)
	if err != nil {
		return err
	}

	exists, err := i.CheckStockAvailability(productDetails)
	if err != nil {
		return err
	}

	if !exists {
		return ErrStockNotAvailable
	}

	totalPrice, err := i.TotalPrice(p.ProductDetails)
	if err != nil {
		return err
	}

	if totalPrice > p.Paid {
		return ErrTotalPriceNotMatch
	}

	change := p.Paid - totalPrice

	if *p.Change != change {
		return ErrChangeNotMatch
	}

	err = i.orderRepository.Create(p)

	if err != nil {
		return err
	}

	return nil
}
