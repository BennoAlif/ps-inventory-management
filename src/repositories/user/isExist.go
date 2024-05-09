package userrepository

import (
	"log"
	"strconv"
	"strings"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sUserRepository) IsExists(filters *entities.ParamsCreateUser) (bool, error) {
	var sb strings.Builder
	var params []interface{}
	var conditions []string

	sb.WriteString("SELECT EXISTS (SELECT 1 FROM users WHERE ")

	if filters.ID != 0 {
		params = append(params, filters.ID)
		conditions = append(conditions, "id = $"+strconv.FormatInt(int64(len(params)), 10))
	}
	if filters.Name != "" {
		params = append(params, filters.Name)
		conditions = append(conditions, "name = $"+strconv.FormatInt(int64(len(params)), 10))
	}
	if filters.PhoneNumber != "" {
		params = append(params, filters.PhoneNumber)
		conditions = append(conditions, "phone_number = $"+strconv.FormatInt(int64(len(params)), 10))
	}

	sb.WriteString(strings.Join(conditions, " AND "))
	sb.WriteString(")")

	var exists bool
	err := i.DB.QueryRow(sb.String(), params...).Scan(&exists)

	if err != nil {
		log.Printf("Error checking if user exists: %s", err)
		return false, err
	}

	return exists, nil
}
