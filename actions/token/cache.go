package token

import (
	"auth-service/actions/login"
	"auth-service/server"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func checkInCache(server *server.APIServer, jwtInfo *jwt.MapClaims) (string, error) {
	tokenInfo := *jwtInfo

	cachekey := fmt.Sprintf("%s-%s", login.CacheKey, tokenInfo["preferred_username"])

	value, err := server.Cache().Get(cachekey)
	if err != nil {
		return "", errors.New("error: no token in cache for user")
	}

	return value, nil
}

func storeTokenToCache(server *server.APIServer, kresp *KeycloakResp, jwtInfo *jwt.MapClaims, jwt string) {

	tokenInfo := *jwtInfo

	expire := kresp.Expire - time.Now().Unix()

	cachekey := fmt.Sprintf("%s-%s", login.CacheKey, tokenInfo["preferred_username"])

	server.Cache().Set(cachekey, jwt, time.Duration(expire*int64(time.Second)))
}
