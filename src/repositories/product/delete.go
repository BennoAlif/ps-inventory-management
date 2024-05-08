package productrepository

import "log"

func (i *sProductRepository) Delete(productId *string) error {
	_, err := i.DB.Exec("DELETE FROM products WHERE id = $1;", productId)

	if err != nil {
		log.Printf("Error deleting product: %s", err)
		return err
	}

	return nil
}
