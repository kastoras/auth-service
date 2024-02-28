package master_realm

import (
	"auth-service/helpers"
	"fmt"
)

func GetManagmentUrl(action string) string {

	url := helpers.GetEnvParam("KEYCKLOACK_ADMIN_MANAGMENT_HOST", "")
	realm := helpers.GetEnvParam("KEYCLOAK_APP_REALM", "")

	managmentUrl := fmt.Sprintf("%s/%s/%s", url, realm, action)

	return managmentUrl
}
