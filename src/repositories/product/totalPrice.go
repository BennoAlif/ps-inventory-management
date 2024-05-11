package productrepository

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (pr *sProductRepository) TotalPrice(details []entities.ProductDetails) (int, error) {
	priceMap := make(map[int]int)

	productIDMap := make(map[int]struct{})
	for _, detail := range details {
		id, err := strconv.Atoi(detail.ProductID)
		if err != nil {
			return 0, fmt.Errorf("invalid product ID: %s", detail.ProductID)
		}
		productIDMap[id] = struct{}{}
	}

	var productIDs []int
	for id := range productIDMap {
		productIDs = append(productIDs, id)
	}

	query := "SELECT id, price FROM products WHERE id IN ("
	params := []interface{}{}
	for _, id := range productIDs {
		query += "$" + strconv.Itoa(len(params)+1) + ","
		params = append(params, id)
	}
	query = strings.TrimSuffix(query, ",") + ")"

	rows, err := pr.DB.Query(query, params...)
	if err != nil {
		log.Printf("Error getting product prices: %s", err)
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, price int
		err := rows.Scan(&id, &price)
		if err != nil {
			log.Printf("Error scanning product price: %s", err)
			return 0, err
		}
		priceMap[id] = price
	}

	var totalPrice int
	for _, detail := range details {
		id, err := strconv.Atoi(detail.ProductID)
		if err != nil {
			return 0, fmt.Errorf("invalid product ID: %s", detail.ProductID)
		}
		price, ok := priceMap[id]
		if !ok {
			return 0, fmt.Errorf("price not found for product ID: %d", id)
		}
		totalPrice += price * detail.Quantity
	}

	return totalPrice, nil
}
