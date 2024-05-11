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
	FindMany(*entities.SearchOrderFilter) ([]*entities.ParamsCustomerCheckout, error)
}

func New(db *sql.DB) OrderRepository {
	return &sOrderRepository{DB: db}
}
