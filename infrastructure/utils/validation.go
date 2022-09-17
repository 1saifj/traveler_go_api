package utils

import "strings"

func ValidateEmail(email string) bool {
	if strings.Contains(email, "@") {
		return true
	}
	return false
}
