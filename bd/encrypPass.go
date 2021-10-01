package bd

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(pass string) (string, error) {
	salt := 8

	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), salt)
	return string(bytes), err
}
