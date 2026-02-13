package pkg

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string, workFactor int) (string, error) {
	byteHashPass, err := bcrypt.GenerateFromPassword([]byte(password), workFactor)
	if err != nil {
		log.Println("加密密码失败!" + err.Error())
		return "", err
	}
	hashPass := string(byteHashPass)
	return hashPass, nil
}
func ComparePassword(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
