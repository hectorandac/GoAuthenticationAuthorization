package services

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/hectorandac/AuthenticationAuthorization/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetClaim(req *http.Request) (models.AuthTokenClaims, error) {
	token := req.Header.Get("Authorization")

	if len(token) > 0 {
		claims := models.AuthTokenClaims{}
		key := func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte("SecretCode"), nil
		}

		_, err := jwt.ParseWithClaims(token, &claims, key)
		if err != nil {
			return models.AuthTokenClaims{}, err
		}

		return claims, nil
	}

	return models.AuthTokenClaims{}, errors.New("no token provided")
}
