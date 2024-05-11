package customerrepository

import (
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sCustomerRepository) FindMany(filters *entities.ParamsCustomer) ([]*entities.Customer, error) {
	query := "SELECT id, name, phone_number FROM customers WHERE 1=1 "
	params := []interface{}{}

	n := (&entities.ParamsCustomer{})

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

		if filters.PhoneNumber != "" {
			query += "AND phone_number ILIKE '+' || $" + strconv.Itoa(len(params)+1) + " || '%'"
			params = append(params, filters.PhoneNumber)
		}

		if len(conditions) > 0 {
			query += " AND "
		}
		query += strings.Join(conditions, " AND ")
	}
	rows, err := i.DB.Query(query, params...)
	if err != nil {
		log.Printf("Error finding customer: %s", err)
		return nil, err
	}
	defer rows.Close()

	customers := make([]*entities.Customer, 0)
	for rows.Next() {
		c := new(entities.Customer)
		err := rows.Scan(&c.ID, &c.Name, &c.PhoneNumber)
		if err != nil {
			log.Printf("Error scanning customer: %s", err)
			return nil, err
		}
		customers = append(customers, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}
