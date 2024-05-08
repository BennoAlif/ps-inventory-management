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

func (i *sUserRepository) FindByPhoneNumber(phone_number *string) (*entities.User, error) {
	row := i.DB.QueryRow("SELECT id, name, phone_number, role, password FROM users WHERE phone_number = $1", phone_number)

	var user entities.User
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Role, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("o user found with email: %s", phone_number)
			return nil, nil // Return nil for both user and error
		}
		log.Printf("Error scanning user by phone number: %s", err)
		return nil, err
	}

	return &user, nil
}
