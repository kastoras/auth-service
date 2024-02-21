package helpers

import (
	"fmt"
)

func BuildKeyclaockAPIUrl(keycloackHost string, endpoint string) string {
	host := fmt.Sprintf("%s/%s", keycloackHost, endpoint)
	return host
}
