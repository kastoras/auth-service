package login

type KloginPayload struct {
	clientID     string
	username     string
	password     string
	grandType    string
	clientSecret string
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
