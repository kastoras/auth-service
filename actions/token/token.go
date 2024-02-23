package token

import (
	"github.com/golang-jwt/jwt/v5"
)

func (tc *TokenController) checkTokenValidy(jwtInfo *jwt.MapClaims, requestJWT string) error {

	token, err := tc.checkInCache(jwtInfo)
	if err == nil && tc.compareTokens(token, requestJWT) {
		return nil
	}

	err = tc.IsTokenActive(jwtInfo, requestJWT)
	if err != nil {
		return err
	}

	return nil
}

func (tc *TokenController) compareTokens(cachedToken string, requestJWT string) bool {
	return cachedToken == requestJWT
}
