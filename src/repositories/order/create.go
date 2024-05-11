package orderepository

import (
	"log"
	"sort"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sOrderRepository) Create(params *entities.ParamsCustomerCheckout) error {
	stmt1, err := i.DB.Prepare(`INSERT INTO orders (customer_id, paid, "change") VALUES ($1, $2, $3) RETURNING id`)
	if err != nil {
		log.Printf("Error preparing order insert: %s", err)
		return err
	}
	defer stmt1.Close()

	stmt2, err := i.DB.Prepare(`INSERT INTO order_details (order_id, product_id, quantity) VALUES ($1, $2, $3)`)
	if err != nil {
		log.Printf("Error preparing order detail insert: %s", err)
		return err
	}
	defer stmt2.Close()

	stmt3, err := i.DB.Prepare(`UPDATE products SET stock = stock - LEAST(stock, $1) WHERE id = $2`)
	if err != nil {
		log.Printf("Error preparing product update: %s", err)
		return err
	}
	defer stmt3.Close()

	var id int
	err = stmt1.QueryRow(params.CustomerID, params.Paid, *params.Change).Scan(&id)
	if err != nil {
		log.Printf("Error inserting order: %s", err)
		return err
	}

	// Sort product details by ProductID
	sort.Slice(params.ProductDetails, func(i, j int) bool {
		return params.ProductDetails[i].ProductID < params.ProductDetails[j].ProductID
	})

	// Then proceed with your existing code
	for _, detail := range params.ProductDetails {
		_, err = stmt2.Exec(id, detail.ProductID, detail.Quantity)
		if err != nil {
			log.Printf("Error inserting order detail: %s", err)
			return err
		}
	}

	for _, detail := range params.ProductDetails {
		_, err = stmt3.Exec(detail.Quantity, detail.ProductID)
		if err != nil {
			log.Printf("Error decreasing product quantity: %s", err)
			return err
		}
	}

	return nil
}
