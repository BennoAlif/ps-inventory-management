package productrepository

import (
	"log"
	"strconv"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sProductRepository) IsExists(filters *entities.ProductSearchFilter) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM products WHERE 1=1 "
	params := []interface{}{}

	if filters.ID != "" {
		query += "AND id = $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.ID)
	}

	query += ")"

	var exists bool
	err := i.DB.QueryRow(query, params...).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if cat exists: %s", err)
		return false, err
	}

	return exists, nil
}
