package helpers

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func GetJWTFromHeader(authHeader string) (string, error) {

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 {
		return "", errors.New("error: invalid Auth Header")
	}

	return parts[1], nil
}

func Parse(jwtToken string) (jwt.MapClaims, error) {

	token, _, err := new(jwt.Parser).ParseUnverified(jwtToken, jwt.MapClaims{})
	if err != nil {
		return jwt.MapClaims{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jwt.MapClaims{}, err
	}

	return claims, nil
}
