package productusecase

import (
	"time"

	"github.com/BennoAlif/ps-cats-social/src/entities"
	"github.com/google/uuid"
)

type (
	ResultCreate struct {
		ID        string `json:"id"`
		CreatedAt string `json:"createdAt"`
	}
)

func (i *sProductUsecase) Create(p *entities.ParamsCreateProduct) (*ResultCreate, error) {
	newID := uuid.New()
	data, err := i.productRepository.Create(&entities.ParamsCreateProduct{
		ID:          newID,
		Name:        p.Name,
		Sku:         p.Sku,
		Category:    p.Category,
		ImageUrl:    p.ImageUrl,
		Notes:       p.Notes,
		Price:       p.Price,
		Stock:       p.Stock,
		Location:    p.Location,
		IsAvailable: p.IsAvailable,
	})

	if err != nil {
		return nil, err
	}

	return &ResultCreate{
		ID:        data.ID,
		CreatedAt: data.CreatedAt.Format(time.RFC3339),
	}, nil

}
