package types

type AccessTokenRequest struct {
	AppID     string `query:"appid"`
	Secret    string `query:"secret"`
	GrantType string `query:"grant_type"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	*Error
}

type StableAccessTokenRequest struct {
	AppID        string `json:"appid"`
	Secret       string `json:"secret"`
	GrantType    string `json:"grant_type"`
	ForceRefresh bool   `json:"force_refresh"`
}

type StableAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	*Error
}
