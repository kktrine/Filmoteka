package auth

import (
	"crypto/rand"
	"encoding/hex"
	"filmoteka_server/internal/repository"

	//"filmoteka_server/internal/repository"
	"filmoteka_server/models"

	"github.com/go-openapi/errors"
)

var tokens = make(map[string]string)

func CheckAdminToken(token string) (*models.Principal, error) {
	role, ok := tokens[token]
	if ok {
		if role == "isAdmin" {
			return &models.Principal{
				Name:  "",
				Roles: []string{role},
			}, nil
		}
		return nil, errors.New(401, "this user hasn't permissions")
	}
	return nil, errors.New(401, "invalid token")
}

func CheckUserToken(token string) (*models.Principal, error) {
	role, ok := tokens[token]
	if ok {
		if role == "isUser" {
			return &models.Principal{
				Name:  "",
				Roles: []string{role},
			}, nil
		}
		return nil, errors.New(401, "this user doesn't have permissions")
	}
	return nil, errors.New(401, "invalid token")

}

func Login(username, password string, storage repository.Repository) (string, error) {
	role, err := storage.Login(username, password)
	if err != nil {
		return "", err
	}
	token, err := generateToken()
	if err != nil {
		return "", err
	}
	for _, ok := tokens[token]; ok; {
		token, err = generateToken()
		if err != nil {
			return "", err
		}
	}
	tokens[token] = role
	return token, nil
}

func generateToken() (string, error) {
	b := make([]byte, 15)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func Logout(token string) {
	delete(tokens, token)
}
