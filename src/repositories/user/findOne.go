package userrepository

import (
	"bytes"
	"database/sql"
	"log"
	"strconv"
	"strings"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sUserRepository) FindOne(filters *entities.ParamsCreateUser) (*entities.User, error) {
	var query bytes.Buffer
	query.WriteString("SELECT id, name, phone_number, password FROM users WHERE ")

	params := []interface{}{}
	conditions := []string{}

	addCondition := func(condition string, param interface{}) {
		conditions = append(conditions, condition+" = $"+strconv.Itoa(len(params)+1))
		params = append(params, param)
	}

	if filters.ID != 0 {
		addCondition("id", filters.ID)
	}
	if filters.Name != "" {
		addCondition("name", filters.Name)
	}
	if filters.PhoneNumber != "" {
		addCondition("phone_number", filters.PhoneNumber)
	}

	query.WriteString(strings.Join(conditions, " AND "))
	query.WriteString(" LIMIT 1")

	row := i.DB.QueryRow(query.String(), params...)

	var user entities.User
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password)

	if err != nil {
		log.Printf("Error find user: %s", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
