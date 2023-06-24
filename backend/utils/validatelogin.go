package utils

import (
	"backend/dao"

	"golang.org/x/crypto/bcrypt"
)

func ValidateLogin(username, password string) (*dao.APIUser, error) {
	user, hashedPassword, err := dao.GetUserByUsername(username, false)
	if err != nil {
		return &dao.APIUser{}, err
	}

	// Check whether the password is correct.
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return &dao.APIUser{}, err
	}

	return user, nil
}
