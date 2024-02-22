package token

import (
	"auth-service/server"

	"github.com/golang-jwt/jwt/v5"
)

func checkTokenValidy(server *server.APIServer, jwtInfo *jwt.MapClaims, requestJWT string) error {

	token, err := checkInCache(server, jwtInfo)
	if err == nil && compareTokens(token, requestJWT) {
		return nil
	}

	err = IsTokenActive(*server, jwtInfo, requestJWT)
	if err != nil {
		return err
	}

	return nil
}

func compareTokens(cachedToken string, requestJWT string) bool {
	return cachedToken == requestJWT
}
