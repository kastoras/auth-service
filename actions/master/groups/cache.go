package master_groups

import (
	"encoding/json"
	"time"
)

func (mgc *MasterGroupsController) storeToCache(groups *[]RealmGroup) error {

	jsonData, err := json.Marshal(&groups)
	if err != nil {
		return err
	}

	mgc.server.Cache().Set(cacheKey, string(jsonData), time.Duration(cacheExpiration*int(time.Second)))

	return nil
}

func (mgc *MasterGroupsController) getFromCache() ([]RealmGroup, error) {

	value, err := mgc.server.Cache().Get(cacheKey)
	if err != nil {
		return []RealmGroup{}, err
	}

	var groupsResp []RealmGroup
	err = json.Unmarshal([]byte(value), &groupsResp)
	if err != nil {
		return []RealmGroup{}, err
	}

	return groupsResp, nil
}
