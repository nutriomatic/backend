package utils

import "golang.org/x/crypto/bcrypt"

func ValidatePassword(hashed_pass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed_pass), []byte(pass)) == nil
}

func ValidateLengthPassword(password string) bool {
	lengthConstraint := len(password) >= 8

	return lengthConstraint
}
