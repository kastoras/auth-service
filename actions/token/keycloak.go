package token

import (
	"auth-service/helpers"
	"auth-service/server"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var (
	endpoint = "protocol/openid-connect/token/introspect"
)

func IsTokenActive(server server.APIServer, jwtTokenInfo *jwt.MapClaims, jwtToken string) error {

	keycloack := server.KCClient()

	kpayload := &KeycloakTokenIntrospectPayload{
		ClientID:     keycloack.ClientID,
		ClientSecret: keycloack.ClientSecret,
		Token:        jwtToken,
	}

	encodedFormData, err := helpers.EncodeForKeycloakRequest(*kpayload)
	if err != nil {
		return err
	}

	host := helpers.BuildKeyclaockAPIUrl(server.KCClient().Host, endpoint)

	req, err := http.NewRequest("POST", host, strings.NewReader(encodedFormData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := server.Client().Do(req)
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

	go storeTokenToCache(&server, &kResp, jwtTokenInfo, jwtToken)

	return nil
}
