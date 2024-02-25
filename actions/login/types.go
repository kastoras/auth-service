package login

type KloginPayload struct {
	ClientID     string `json:"client_id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	GrandType    string `json:"grant_type"`
	ClientSecret string `json:"client_secret"`
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type KLoginResp struct {
	AccessToken string  `json:"access_token"`
	Expiration  float64 `json:"expires_in"`
}

type LoginResp struct {
	AccessToken string `json:"access_token"`
}
