package master_realm

import (
	"auth-service/helpers"
	"fmt"
)

func GetManagmentUrl(action string, realm string) string {

	url := helpers.GetEnvParam("KEYCKLOACK_ADMIN_MANAGMENT_HOST", "")

	managmentUrl := fmt.Sprintf("%s/%s/%s", url, realm, action)

	return managmentUrl
}
