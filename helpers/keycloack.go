package helpers

import (
	"fmt"
	"net/url"
	"reflect"
)

func BuildKeyclaockAPIUrl(keycloackHost string, endpoint string) string {
	host := fmt.Sprintf("%s/%s", keycloackHost, endpoint)
	return host
}

func EncodeForKeycloakRequest(data interface{}) (string, error) {

	values := url.Values{}

	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("input is not a struct")
	}

	v := reflect.ValueOf(data)

	// Iterate over the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		// Get the tag value for "json"
		tag := field.Tag.Get("json")

		// If the tag is not empty, use it as the key name
		if tag != "" {
			values.Add(tag, fmt.Sprintf("%v", fieldValue.Interface()))
		}
	}

	return values.Encode(), nil
}
