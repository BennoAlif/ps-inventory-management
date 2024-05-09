package productrepository

import (
	"log"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sProductRepository) Update(productId *string, p *entities.ParamsUpdateProduct) error {

	_, err := i.DB.Exec("UPDATE products SET name = $2, sku = $3, category = $4, notes = $5, image_url = $6, price = $7, stock = $8, location = $9, is_available = $10 WHERE id = $1",
		productId,
		p.Name,
		p.Sku,
		p.Category,
		p.Notes,
		p.ImageUrl,
		p.Price,
		p.Stock,
		p.Location,
		p.IsAvailable,
	)

	if err != nil {
		log.Printf("Error updating product: %s", err)
		return err
	}

	return nil
}
