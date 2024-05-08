package productusecase

import (
	"time"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sProductUsecase) Update(productId *string, p *entities.ParamsUpdateProduct) (*ResultCreate, error) {
	filters := entities.ProductSearchFilter{
		ID: *productId,
	}

	product, _ := i.productRepository.IsExists(&filters)

	if !product {
		return nil, ErrProductNotFound
	}

	data, err := i.productRepository.Update(productId,
		&entities.ParamsUpdateProduct{
			Name:        p.Name,
			Sku:         p.Sku,
			Category:    p.Category,
			Notes:       p.Notes,
			ImageUrl:    p.ImageUrl,
			Price:       p.Price,
			Stock:       p.Stock,
			Location:    p.Location,
			IsAvailable: p.IsAvailable,
		},
	)

	if err != nil {
		return nil, err
	}

	return &ResultCreate{
		ID:        data.ID,
		CreatedAt: data.CreatedAt.Format(time.RFC3339),
	}, nil

}
