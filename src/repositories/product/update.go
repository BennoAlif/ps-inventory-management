package productrepository

import (
	"log"
	"time"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sProductRepository) Update(productId *string, p *entities.ParamsUpdateProduct) (*entities.CreateProduct, error) {
	var id string
	var createdAt time.Time

	err := i.DB.QueryRow("UPDATE products SET name = $2, sku = $3, category = $4, image_url = $5, notes = $6, price = $7, location = $8, is_available = $9 WHERE id = $1 RETURNING id, created_at",
		productId,
		p.Name,
		p.Sku,
		p.Category,
		p.ImageUrl,
		p.Notes,
		p.ImageUrl,
		p.Location,
		p.IsAvailable,
	).Scan(&id, &createdAt)

	if err != nil {
		log.Printf("Error updating product: %s", err)
		return nil, err
	}

	product := &entities.CreateProduct{
		ID:        id,
		CreatedAt: createdAt,
	}

	return product, nil
}
