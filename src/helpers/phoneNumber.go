package helpers

import (
	"errors"
	"regexp"
)

var (
	ErrBadFormatPhoneNumber = errors.New("Format Phone tidak valid")
)

func IsValidPhoneNumber(phoneNumber string) error {
	phoneNumberRegex := `^\+[0-9]{1,4}-?[0-9]{1,15}$`
	match, _ := regexp.MatchString(phoneNumberRegex, phoneNumber)
	if match {
		return nil
	}

	return ErrBadFormatPhoneNumber
}
