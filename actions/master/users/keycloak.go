package master_users

import (
	master_realm "auth-service/helpers/masterrealm"
	master_token "auth-service/helpers/masterrealm/token"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var (
	getUsersEntpoint      = "users"
	getCountUsersEndpoint = "users/count"
)

func (muc *MasterUsersController) List(pagination Pagination) ([]User, error) {
	host := master_realm.GetManagmentUrl(getUsersEntpoint, muc.server.KCClient().Realm)

	host = addPagination(host, pagination.Offset, pagination.Limit)

	resp, err := muc.executeAPIRequest(host)
	if err != nil {
		fmt.Printf("Error : Failed to comunicate to keycloak %v \n", err)
	}

	var usersResp []User
	err = json.NewDecoder(resp.Body).Decode(&usersResp)
	if err != nil {
		fmt.Printf("Error : Failed to parse keycloak users response %v \n", err)
		return []User{}, errors.New("login failed")
	}

	return usersResp, nil
}

func (muc *MasterUsersController) TotalUsers() (int, error) {
	host := master_realm.GetManagmentUrl(getCountUsersEndpoint, muc.server.KCClient().Realm)

	resp, err := muc.executeAPIRequest(host)
	if err != nil {
		fmt.Printf("Error : Failed to comunicate to keycloak %v \n", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return -1, errors.New("error reading response body")
	}

	numString := string(body)
	intValue, err := strconv.Atoi(numString)
	if err != nil {
		// Handle error
		fmt.Println("Error parsing integer:", err)
		return -1, errors.New("error parsing integer")
	}

	return intValue, nil
}

func (muc *MasterUsersController) executeAPIRequest(host string) (*http.Response, error) {
	req, err := http.NewRequest("GET", host, nil)
	if err != nil {
		return nil, err
	}

	masterToken, err := master_token.Get(muc.server)
	if err != nil {
		return nil, err
	}

	bearerToken := fmt.Sprintf("Bearer %s", masterToken)

	req.Header.Add("Authorization", bearerToken)
	resp, err := muc.server.Client().Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("login failed")
	}

	return resp, nil
}

func addPagination(url string, start, limit int) string {
	return fmt.Sprintf("%s?max=%d&first=%d", url, limit, start)
}
