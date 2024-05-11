package productrepository

import (
	"log"
	"strconv"
)

func (i *sProductRepository) IsExistsMany(filters []*string) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM products WHERE 1=1 AND is_available = true "
	params := []interface{}{}

	if len(filters) > 0 {
		query += "AND id IN ("
		for _, filter := range filters {
			query += "$" + strconv.Itoa(len(params)+1) + ","
			params = append(params, *filter)
		}
		query = query[:len(query)-1] + ")"
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
