package utils

import (
	"Houses/internal/model"
	"golang.org/x/crypto/bcrypt"
)

const MinPasswordLength = 4

func GenPassHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func CheckUserLogin(login model.UserLogin) error {
	if len(login.Password) < MinPasswordLength {
		return ErrPasswordTooShort
	}

	if login.Email == "" {
		return ErrEmptyEmail
	}

	return nil
}
