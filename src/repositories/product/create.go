package productrepository

import (
	"log"
	"time"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sProductRepository) Create(p *entities.ParamsCreateProduct) (*entities.CreateProduct, error) {
	var id string
	var createdAt time.Time
	err := i.DB.QueryRow("INSERT INTO products (name, sku, category, image_url, notes, price, stock, location, is_available) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, created_at",
		p.Name,
		p.Sku,
		p.Category,
		p.ImageUrl,
		p.Notes,
		p.Price,
		p.Stock,
		p.Location,
		p.IsAvailable,
	).Scan(&id, &createdAt)

	if err != nil {
		log.Printf("Error creating product: %s", err)
		return nil, err
	}

	product := &entities.CreateProduct{
		ID:        id,
		CreatedAt: createdAt,
	}

	return product, nil
}
