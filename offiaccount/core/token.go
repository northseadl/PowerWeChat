package core

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/ArtisanCloud/PowerWeChat/pro/offiaccount/types"
	"github.com/ArtisanCloud/PowerWeChat/pro/pkg/cachex"
	"github.com/artisancloud/httphelper"
)

const (
	TokenUrl       = "https://api.weixin.qq.com/cgi-bin/token"
	StabelTokenUrl = "https://api.weixin.qq.com/cgi-bin/component/stable_token"
)

const tokenRefreshGracePeriod = 60 * time.Second

type TokenSource struct {
	AccessToken *Token
	c           *AuthConfig
	h           httphelper.Helper
	cache       cachex.Cache
	mu          sync.Mutex
}

type AuthConfig struct {
	AppID          string
	Secret         string
	UseStableToken bool
}

func newTokenSource(h httphelper.Helper, c AuthConfig, cache cachex.Cache) *TokenSource {
	return &TokenSource{
		c:     &c,
		h:     h,
		cache: cache,
	}
}

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

func (t *Token) String() string {
	return t.AccessToken
}

func (t *TokenSource) Valid() bool {
	if t.AccessToken == nil {
		return false
	}
	if t.AccessToken.ExpiresAt < time.Now().Unix()+int64(tokenRefreshGracePeriod) {
		return false
	}
	return true
}

func (t *TokenSource) GetAccessToken() (string, error) {
	if t.Valid() {
		return t.AccessToken.AccessToken, nil
	}
	mashaledToken, err := t.cache.Get(formartCacheKey(t.c.AppID))
	if err == nil {
		token, err := unmarshalToken(mashaledToken)
		if err != nil {
			return "", err
		}
		t.AccessToken = token
		return token.AccessToken, nil
	}
	token, err := t.Refresh()
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}

// Refresh
func (t *TokenSource) Refresh() (*Token, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.Valid() {
		return t.AccessToken, nil
	}

	var token *Token
	var err error
	if t.c.UseStableToken {
		token, err = t.GetStableAccessTokenFromServer()
	} else {
		token, err = t.GetAccessTokenFromServer()
	}
	if err != nil {
		return nil, err
	}

	t.AccessToken = token
	mashaledToken, err := mashalToken(token)
	if err != nil {
		return nil, err
	}
	err = t.cache.Set(formartCacheKey(t.c.AppID), mashaledToken, time.Duration(token.ExpiresAt-time.Now().Unix())*time.Second)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// GetAccessTokenFromServer
func (t *TokenSource) GetAccessTokenFromServer() (*Token, error) {
	now := time.Now().Unix()
	var result types.AccessTokenResponse
	err := t.h.Df().Method("POST").Url(TokenUrl).
		BindQuery(&types.AccessTokenRequest{
			AppID:     t.c.AppID,
			Secret:    t.c.Secret,
			GrantType: "client_credential",
		}).
		Result(&result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, result.Error
	}
	token := &Token{
		AccessToken: result.AccessToken,
		ExpiresAt:   now + result.ExpiresIn,
	}
	return token, nil
}

// GetStableAccessTokenFromServer
func (t *TokenSource) GetStableAccessTokenFromServer() (*Token, error) {
	now := time.Now().Unix()
	var result types.StableAccessTokenResponse
	err := t.h.Df().Method("POST").Url(StabelTokenUrl).
		BindQuery(types.StableAccessTokenRequest{
			AppID:     t.c.AppID,
			Secret:    t.c.Secret,
			GrantType: "client_credential",
		}).
		Result(&result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, result.Error
	}
	token := &Token{
		AccessToken: result.AccessToken,
		ExpiresAt:   now + result.ExpiresIn,
	}
	return token, nil
}

func formartCacheKey(appID string) string {
	return "powerwechat:offiaccount:access_token:" + appID
}

func mashalToken(token *Token) (string, error) {
	data, err := json.Marshal(token)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func unmarshalToken(data string) (*Token, error) {
	var token Token
	err := json.Unmarshal([]byte(data), &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
