package userrepository

import (
	"log"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

type (
	ParamsCreateUser struct {
		Name        string
		PhoneNumber string
		Password    string
	}
)

func (i *sUserRepository) Create(p *ParamsCreateUser) (*entities.User, error) {
	var id int64
	role := "STAFF"
	err := i.DB.QueryRow("INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id", p.Name, p.PhoneNumber, role, p.Password).Scan(&id)
	if err != nil {
		log.Printf("Error inserting user: %s", err)
		return nil, err
	}

	user := &entities.User{
		ID:          id,
		Name:        p.Name,
		PhoneNumber: p.PhoneNumber,
		Role:        role,
	}

	return user, nil
}
