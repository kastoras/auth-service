package master_token

type KMasterloginPayload struct {
	ClientID     string `json:"client_id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	GrandType    string `json:"grant_type"`
	ClientSecret string `json:"client_secret"`
}

type KMasterLoginResp struct {
	AccessToken string  `json:"access_token"`
	Expiration  float64 `json:"expires_in"`
}

type MasterTokenController struct{}
