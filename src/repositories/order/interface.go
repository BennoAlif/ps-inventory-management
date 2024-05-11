package orderepository

import (
	"database/sql"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

type sOrderRepository struct {
	DB *sql.DB
}

type OrderRepository interface {
	Create(*entities.ParamsCustomerCheckout) error
	// FindMany(*entities.OrderSearchFilter) ([]*entities.Order, error)
}

func New(db *sql.DB) OrderRepository {
	return &sOrderRepository{DB: db}
}
