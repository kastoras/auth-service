package master_groups

import (
	"encoding/json"
	"fmt"
	"time"
)

func (mgc *MasterGroupsController) storeToCache(groups *[]RealmGroup) error {

	cacheKey := mgc.getCacheKey()

	jsonData, err := json.Marshal(&groups)
	if err != nil {
		return err
	}

	mgc.server.Cache().Set(cacheKey, string(jsonData), time.Duration(cacheExpiration*int(time.Second)))

	return nil
}

func (mgc *MasterGroupsController) getFromCache() ([]RealmGroup, error) {

	cacheKey := mgc.getCacheKey()

	value, err := mgc.server.Cache().Get(cacheKey)
	if err != nil {
		return []RealmGroup{}, err
	}

	var groupsResp []RealmGroup
	err = json.Unmarshal([]byte(value), &groupsResp)

	return groupsResp, err
}

func (mgc *MasterGroupsController) getCacheKey() string {
	return fmt.Sprintf("%s-groups", mgc.server.KCClient().Realm)
}
