package master_token

import "auth-service/server"

func Get(server *server.APIServer) (string, error) {

	token, err := server.Cache().Get(cacheKey)
	if err == nil {
		return token, nil
	}

	kresp, err := keycloackMasterLogin(server)
	if err != nil {
		return "", err
	}

	return kresp.AccessToken, err
}
