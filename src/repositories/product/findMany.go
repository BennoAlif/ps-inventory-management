package productrepository

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sProductRepository) FindMany(filters *entities.ProductSearchFilter) ([]*entities.Product, error) {
	query := "SELECT id, name, sku, category, image_url, stock, notes, price, location, is_available, created_at FROM products WHERE 1=1 "
	params := []interface{}{}

	n := (&entities.ProductSearchFilter{})

	if !reflect.DeepEqual(filters, n) {
		conditions := []string{}

		if filters.ID != "" {
			conditions = append(conditions, "id = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.ID)
		}
		if filters.Name != "" {
			query += "AND name ILIKE '%' || $" + strconv.Itoa(len(params)+1) + " || '%'"
			params = append(params, filters.Name)
		}
		if filters.IsAvailable != nil && (filters.IsAvailable == true || filters.IsAvailable == false) {
			query += "AND is_available = $" + strconv.Itoa(len(params)+1)
			params = append(params, filters.IsAvailable)
		}
		if filters.Category != "" {
			query += "AND category = $" + strconv.Itoa(len(params)+1)
			params = append(params, filters.Category)
		}
		if filters.SKU != "" {
			query += "AND sku = $" + strconv.Itoa(len(params)+1)
			params = append(params, filters.SKU)
		}
		if filters.InStock != nil && filters.InStock.(bool) {
			query += "AND stock > 0"
		} else if filters.InStock != nil && !filters.InStock.(bool) {
			query += "AND stock = 0"
		}

		if len(conditions) > 0 {
			query += " AND "
		}
		query += strings.Join(conditions, " AND ")
	}

	if filters.Limit == 0 {
		filters.Limit = 5
	}

	sortingQuery := " ORDER BY"
	sortingStr := []string{}

	if filters.CreatedAt == "asc" {
		sortingStr = append(sortingStr, " created_at ASC")
	} else if filters.CreatedAt == "desc" {
		sortingStr = append(sortingStr, " created_at DESC")
	}

	if filters.Price == "asc" {
		sortingStr = append(sortingStr, " price ASC")
	} else if filters.Price == "desc" {
		sortingStr = append(sortingStr, " price DESC")
	}

	if len(sortingStr) > 0 {
		sortingQuery += strings.Join(sortingStr, ", ")
		query += sortingQuery
	}

	query += " LIMIT $" + strconv.Itoa(len(params)+1)
	params = append(params, filters.Limit)

	if filters.Offset == 0 {
		filters.Offset = 0
	} else {
		query += " OFFSET $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.Offset)
	}

	fmt.Println(query)

	rows, err := i.DB.Query(query, params...)
	if err != nil {
		log.Printf("Error finding product: %s", err)
		return nil, err
	}
	defer rows.Close()

	products := make([]*entities.Product, 0)
	for rows.Next() {
		c := new(entities.Product)
		err := rows.Scan(&c.ID, &c.Name, &c.SKU, &c.Category, &c.ImageUrl, &c.Stock, &c.Notes, &c.Price, &c.Location, &c.IsAvailable, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
