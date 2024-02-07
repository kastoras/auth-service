package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) login(payload *KloginPayload) (*KLoginResp, error) {

	formData := url.Values{
		"client_id":     {payload.clientID},
		"username":      {payload.username},
		"password":      {payload.password},
		"grant_type":    {payload.grandType},
		"client_secret": {payload.clientSecret},
	}

	encodedFormData := formData.Encode()

	req, err := http.NewRequest("POST", "http://keycloak-keycloak-1:8080/realms/authentication-api/protocol/openid-connect/token", strings.NewReader(encodedFormData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Something went wrong on keycloack")
	}

	kloginresp := &KLoginResp{}
	json.NewDecoder(resp.Body).Decode(kloginresp)

	return kloginresp, nil
}
