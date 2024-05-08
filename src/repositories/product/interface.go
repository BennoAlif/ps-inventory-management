package productrepository

import (
	"database/sql"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

type sProductRepository struct {
	DB *sql.DB
}

type ProductRepository interface {
	Create(*entities.ParamsCreateProduct) (*entities.CreateProduct, error)
	FindMany(*entities.ProductSearchFilter) ([]*entities.Product, error)
	IsExists(*entities.ProductSearchFilter) (bool, error)
	Update(*string, *entities.ParamsUpdateProduct) (*entities.CreateProduct, error)
	Delete(*string) error
}

func New(db *sql.DB) ProductRepository {
	return &sProductRepository{DB: db}
}
