package orderepository

import (
	"log"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sOrderRepository) Create(params *entities.ParamsCustomerCheckout) error {
	tx, err := i.DB.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %s", err)
		return err
	}

	var id int

	query := `INSERT INTO orders (customer_id, paid, "change") VALUES ($1, $2, $3) RETURNING id`
	err = tx.QueryRow(query, params.CustomerID, params.Paid, *params.Change).Scan(&id)

	if err != nil {
		log.Printf("Error inserting order: %s", err)
		tx.Rollback()
		return err
	}

	query = `INSERT INTO order_details (order_id, product_id, quantity) VALUES ($1, $2, $3)`
	for _, detail := range params.ProductDetails {
		_, err = tx.Exec(query, id, detail.ProductID, detail.Quantity)
		if err != nil {
			log.Printf("Error inserting order detail: %s", err)
			tx.Rollback()
			return err
		}
	}

	for _, detail := range params.ProductDetails {
		query = `UPDATE products SET stock = stock - $1 WHERE id = $2`
		_, err = tx.Exec(query, detail.Quantity, detail.ProductID)
		if err != nil {
			log.Printf("Error decreasing product quantity: %s", err)
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error committing transaction 4: %s", err)
		tx.Rollback()
		return err
	}

	return nil
}
