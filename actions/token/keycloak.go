package token

import (
	"auth-service/helpers"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var (
	endpoint = "protocol/openid-connect/token/introspect"
)

func (tc *TokenController) IsTokenActive(jwtTokenInfo *jwt.MapClaims, jwtToken string) error {

	keycloack := tc.Server.KCClient()

	kpayload := &KeycloakTokenIntrospectPayload{
		ClientID:     keycloack.ClientID,
		ClientSecret: keycloack.ClientSecret,
		Token:        jwtToken,
	}

	encodedFormData, err := helpers.EncodeForKeycloakRequest(*kpayload)
	if err != nil {
		return err
	}

	host := helpers.BuildKeyclaockAPIUrl(tc.Server.KCClient().Host, endpoint)

	req, err := http.NewRequest("POST", host, strings.NewReader(encodedFormData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := tc.Server.Client().Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("error: introspect failed")
	}

	var kResp KeycloakResp
	err = json.NewDecoder(resp.Body).Decode(&kResp)
	if err != nil {
		return err
	}

	if !kResp.Active {
		return errors.New("fail: inactive token")
	}

	tc.storeTokenToCache(&kResp, jwtTokenInfo, jwtToken)

	return nil
}
