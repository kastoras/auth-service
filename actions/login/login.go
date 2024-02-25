package login

import (
	"auth-service/helpers"
	"auth-service/server"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var (
	endpoint = "protocol/openid-connect/token"
	CacheKey = "access-token"
)

func keycloackLogin(server *server.APIServer, payload LoginPayload) (KLoginResp, error) {

	keycloack := server.KCClient()

	kpayload := &KloginPayload{
		ClientID:     keycloack.ClientID,
		Username:     payload.Username,
		Password:     payload.Password,
		GrandType:    keycloack.ClientGrandType,
		ClientSecret: keycloack.ClientSecret,
	}

	encodedFormData, err := helpers.EncodeForKeycloakRequest(*kpayload)
	if err != nil {
		return KLoginResp{}, err
	}

	host := helpers.BuildKeyclaockAPIUrl(server.KCClient().Host, endpoint)

	req, err := http.NewRequest("POST", host, strings.NewReader(encodedFormData))
	if err != nil {
		fmt.Printf("Error : Failed to create keycloak login request %v \n", err)
		return KLoginResp{}, errors.New("login failed")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := server.Client().Do(req)
	if err != nil {
		fmt.Printf("Error : Failed send keycloak login request %v \n", err)
		return KLoginResp{}, errors.New("login failed")
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error : Failed to login on keycloak status code: %v, error: %v \n", resp.StatusCode, resp)
		return KLoginResp{}, errors.New("login failed")
	}

	var kloginresp KLoginResp
	err = json.NewDecoder(resp.Body).Decode(&kloginresp)
	if err != nil {
		fmt.Printf("Error : Failed to parse keycloak response %v \n", err)
		return KLoginResp{}, errors.New("login failed")
	}

	cacheKey := fmt.Sprintf("%s-%s", CacheKey, kpayload.Username)

	server.Cache().Set(cacheKey, kloginresp.AccessToken, time.Duration(kloginresp.Expiration*float64(time.Second)))

	return kloginresp, nil
}
