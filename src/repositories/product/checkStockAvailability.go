package productrepository

import (
	"fmt"
	"log"
	"strconv"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sProductRepository) CheckStockAvailability(details []*entities.ProductDetails) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM products WHERE 1=1 "
	params := []interface{}{}

	for _, detail := range details {
		_, err := strconv.Atoi(detail.ProductID)
		if err != nil {
			return false, err
		}
		query += fmt.Sprintf("AND id = $%d AND stock >= $%d ", len(params)+1, len(params)+2)
		params = append(params, detail.ProductID, detail.Quantity)
	}

	query += ")"

	var exists bool
	err := i.DB.QueryRow(query, params...).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if product exists: %s", err)
		return false, err
	}

	return exists, nil
}
