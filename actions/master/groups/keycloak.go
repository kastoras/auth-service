package master_groups

import (
	master_realm "auth-service/helpers/masterrealm"
	master_token "auth-service/helpers/masterrealm/token"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	endpoint        = "ui-ext/groups"
	cacheExpiration = 300
)

func (mgc *MasterGroupsController) keycloackGroups() ([]RealmGroup, error) {

	host := master_realm.GetManagmentUrl(endpoint, mgc.server.KCClient().Realm)

	req, err := http.NewRequest("GET", host, nil)
	if err != nil {
		fmt.Printf("Error : Failed to create get groups request %v \n", err)
		return []RealmGroup{}, errors.New("error: failed to create get groups request")
	}

	masterToken, err := master_token.Get(mgc.server)
	if err != nil {
		return []RealmGroup{}, errors.New("error: failed to create admin token")
	}

	bearerToken := fmt.Sprintf("Bearer %s", masterToken)

	req.Header.Add("Authorization", bearerToken)
	resp, err := mgc.server.Client().Do(req)
	if err != nil {
		fmt.Printf("Error : Failed send get groups request %v \n", err)
		return []RealmGroup{}, errors.New("failed send get groups")
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error : Failed to master login on keycloak status code: %v, error: %v \n", resp.StatusCode, resp)
		return []RealmGroup{}, errors.New("login failed")
	}

	var groupsResp []RealmGroup
	err = json.NewDecoder(resp.Body).Decode(&groupsResp)
	if err != nil {
		fmt.Printf("Error : Failed to parse keycloak groups response %v \n", err)
		return []RealmGroup{}, errors.New("login failed")
	}

	mgc.storeToCache(&groupsResp)

	return groupsResp, nil
}
