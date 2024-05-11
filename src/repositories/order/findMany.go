package orderepository

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sOrderRepository) FindMany(filters *entities.SearchOrderFilter) ([]*entities.ParamsCustomerCheckout, error) {
	query := "SELECT id, customer_id, paid, change, created_at FROM orders WHERE 1=1 "
	params := []interface{}{}

	n := (&entities.SearchOrderFilter{})

	if !reflect.DeepEqual(filters, n) {
		conditions := []string{}

		if filters.CustomerID != "" {
			conditions = append(conditions, "customer_id = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.CustomerID)
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

	if len(sortingStr) > 0 {
		sortingQuery += strings.Join(sortingStr, ", ")
		query += sortingQuery
	} else {
		query += " ORDER BY created_at DESC"
	}

	query += " LIMIT $" + strconv.Itoa(len(params)+1)
	params = append(params, filters.Limit)

	query += " OFFSET $" + strconv.Itoa(len(params)+1)
	params = append(params, filters.Offset)

	rows, err := i.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []*entities.ParamsCustomerCheckout{}
	for rows.Next() {
		order := entities.ParamsCustomerCheckout{}
		err := rows.Scan(&order.TransactionID, &order.CustomerID, &order.Paid, &order.Change, &order.CreatedAt)
		if err != nil {
			return nil, err
		}

		detailsQuery := "SELECT product_id, quantity FROM order_details WHERE order_id = $1"
		detailsRows, err := i.DB.Query(detailsQuery, order.TransactionID)
		if err != nil {
			return nil, err
		}
		defer detailsRows.Close()

		orderDetails := []entities.ProductDetails{}
		for detailsRows.Next() {
			detail := entities.ProductDetails{}
			err := detailsRows.Scan(&detail.ProductID, &detail.Quantity)
			if err != nil {
				return nil, err
			}
			orderDetails = append(orderDetails, detail)
		}

		order.ProductDetails = orderDetails
		orders = append(orders, &order)
	}

	return orders, nil
}
