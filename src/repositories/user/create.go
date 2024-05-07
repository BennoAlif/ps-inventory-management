package userrepository

import (
	"log"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

type (
	ParamsCreateUser struct {
		Name     string
		Email    string
		Password string
	}
)

func (i *sUserRepository) Create(p *ParamsCreateUser) (*entities.User, error) {
	var id int64
	err := i.DB.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", p.Name, p.Email, p.Password).Scan(&id)
	if err != nil {
		log.Printf("Error inserting user: %s", err)
		return nil, err
	}

	user := &entities.User{
		ID:    id,
		Name:  p.Name,
		Email: p.Email,
	}

	return user, nil
}
