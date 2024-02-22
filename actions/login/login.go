package login

import (
	"auth-service/helpers"
	"auth-service/server"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	endpoint = "protocol/openid-connect/token"
	cacheKey = "access-token"
)

func keycloackLogin(server *server.APIServer, payload LoginPayload) (KLoginResp, error) {

	keycloack := server.KCClient()

	kpayload := &KloginPayload{
		clientID:     keycloack.ClientID,
		username:     payload.Username,
		password:     payload.Password,
		grandType:    keycloack.ClientGrandType,
		clientSecret: keycloack.ClientSecret,
	}

	formData := url.Values{
		"client_id":     {kpayload.clientID},
		"username":      {kpayload.username},
		"password":      {kpayload.password},
		"grant_type":    {kpayload.grandType},
		"client_secret": {kpayload.clientSecret},
	}

	encodedFormData := formData.Encode()

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

	cacheKey := fmt.Sprintf("%s-%s", cacheKey, kpayload.username)

	server.Cache().Set(cacheKey, kloginresp.AccessToken, time.Duration(kloginresp.Expiration*float64(time.Second)))

	return kloginresp, nil
}
