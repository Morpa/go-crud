package model

import "golang.org/x/crypto/bcrypt"

func (ud *userDomain) EncryptPassword() (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ud.password), 12)
	if err != nil {
		return "", err
	}
	ud.password = string(hashedPassword)
	return string(hashedPassword), nil
}
