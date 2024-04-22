package register

import (
	master_realm "auth-service/helpers/masterrealm"
	master_token "auth-service/helpers/masterrealm/token"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	registerEndpoint  = "users"
	setPasswordAction = "users/%s/reset-password"
)

func (ruc *RegisterUserController) Register(payload RegisterPayload) (int, error) {
	userid, statusCode, err := ruc.registerUser(&payload)
	if err != nil {
		return statusCode, err
	}

	upStatusCode, err := ruc.setUserPassword(userid, payload.Password)
	if err != nil {
		return upStatusCode, err
	}

	return http.StatusCreated, nil
}

func (ruc *RegisterUserController) registerUser(payload *RegisterPayload) (string, int, error) {

	kpayload := &KRegisterPayload{
		Username:        payload.Username,
		Email:           payload.Email,
		EmailVerified:   false, // should be handled by env var or by request
		Firstname:       payload.Firstname,
		Lastname:        payload.Firstname,
		Groups:          []string{"/client"}, // should be handled by env var or by request
		RequiredActions: []string{},          // should be handled by env var or by request
		Enabled:         true,                // should be handled by env var or by request
	}

	jsonPayload, err := json.Marshal(kpayload)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	host := master_realm.GetManagmentUrl(registerEndpoint, ruc.server.KCClient().Realm)

	req, err := http.NewRequest("POST", host, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	masterToken, err := master_token.Get(ruc.server)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	bearerToken := fmt.Sprintf("Bearer %s", masterToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearerToken)
	resp, err := ruc.server.Client().Do(req)
	if err != nil {
		return "", resp.StatusCode, err
	}

	if resp.StatusCode != http.StatusCreated {
		errorMessage := fmt.Sprintf("Error : Failed to register user on keycloak status code: %v, resp %v \n", resp.StatusCode, resp.Body)
		return "", resp.StatusCode, errors.New(errorMessage)
	}

	userID := ruc.getUserId(resp)

	return userID, resp.StatusCode, nil
}

func (ruc *RegisterUserController) getUserId(resp *http.Response) string {
	header := resp.Header.Get("Location")
	splittedPath := strings.Split(header, "/")
	return splittedPath[len(splittedPath)-1]
}

func (ruc *RegisterUserController) setUserPassword(userId string, password string) (int, error) {

	kPassPayload := &KPasswordPayload{
		Temporary: false, // should be handled by env var or by request
		Type:      "password",
		Value:     password,
	}

	userHost := fmt.Sprintf(setPasswordAction, userId)
	host := master_realm.GetManagmentUrl(userHost, ruc.server.KCClient().Realm)

	jsonPayload, err := json.Marshal(kPassPayload)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	req, err := http.NewRequest("PUT", host, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	masterToken, err := master_token.Get(ruc.server)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	bearerToken := fmt.Sprintf("Bearer %s", masterToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearerToken)
	resp, err := ruc.server.Client().Do(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if resp.StatusCode != http.StatusNoContent {
		errorMessage := fmt.Sprintf("Failed to change user password, resp code: %v \n", resp.StatusCode)
		return resp.StatusCode, errors.New(errorMessage)
	}

	return resp.StatusCode, nil
}
