package userrepository

import (
	"database/sql"
	"log"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func (i *sUserRepository) FindByEmail(email *string) (*entities.User, error) {
	row := i.DB.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1", email)

	var user entities.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user found with email: %s", email)
			return nil, nil // Return nil for both user and error
		}
		log.Printf("Error scanning user by email: %s", err)
		return nil, err
	}

	return &user, nil
}
