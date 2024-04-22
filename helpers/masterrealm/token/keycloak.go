package master_token

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
	cacheKey = "master-token"
)

func keycloackMasterLogin(server *server.APIServer) (KMasterLoginResp, error) {

	masterHost := helpers.GetEnvParam("KEYCKLOACK_ADMIN_HOST", "")

	kpayload := &KMasterloginPayload{
		ClientID:     helpers.GetEnvParam("KEYCKLOACK_ADMIN_CLIENT_ID", ""),
		Username:     helpers.GetEnvParam("KEYCKLOACK_ADMIN_USER", ""),
		Password:     helpers.GetEnvParam("KEYCKLOACK_ADMIN_PASS", ""),
		GrandType:    helpers.GetEnvParam("KEYCKLOACK_ADMIN_GRAND_TYPE", ""),
		ClientSecret: helpers.GetEnvParam("KEYCKLOACK_ADMIN_CLIENT_SECRET", ""),
	}

	encodedFormData, err := helpers.EncodeForKeycloakRequest(*kpayload)
	if err != nil {
		return KMasterLoginResp{}, err
	}

	host := helpers.BuildKeyclaockAPIUrl(masterHost, endpoint)

	req, err := http.NewRequest("POST", host, strings.NewReader(encodedFormData))
	if err != nil {
		fmt.Printf("Error : Failed to create keycloak master login request %v \n", err)
		return KMasterLoginResp{}, errors.New("master login failed")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := server.Client().Do(req)
	if err != nil {
		fmt.Printf("Error : Failed send keycloak master login request %v \n", err)
		return KMasterLoginResp{}, errors.New("login failed")
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error : Failed to master login on keycloak status code: %v, error: %v \n", resp.StatusCode, resp)
		return KMasterLoginResp{}, errors.New("login failed")
	}

	var kloginresp KMasterLoginResp
	err = json.NewDecoder(resp.Body).Decode(&kloginresp)
	if err != nil {
		fmt.Printf("Error : Failed to parse keycloak response %v \n", err)
		return KMasterLoginResp{}, errors.New("login failed")
	}

	server.Cache().Set(cacheKey, kloginresp.AccessToken, time.Duration(kloginresp.Expiration*float64(time.Second)))

	return kloginresp, nil
}
