package userrepository

import (
	"database/sql"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

type sUserRepository struct {
	DB *sql.DB
}

type UserRepository interface {
	Create(*ParamsCreateUser) (*entities.User, error)
	FindOne(*entities.ParamsCreateUser) (*entities.User, error)
	IsExists(*entities.ParamsCreateUser) (bool, error)
}

func New(db *sql.DB) UserRepository {
	return &sUserRepository{DB: db}
}
