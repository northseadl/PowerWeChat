package core

import (
	"github.com/ArtisanCloud/PowerWeChat/pro/pkg/cachex"
	"github.com/ArtisanCloud/PowerWeChat/pro/weconfig"
	"github.com/artisancloud/httphelper"
	"github.com/artisancloud/httphelper/client"
)

const endpoint = "https://api.weixin.qq.com/cgi-bin"

type Provider struct {
	Conf *weconfig.Offiaccount
	H    httphelper.Helper
	TS   *TokenSource
}

func NewProvider(c *weconfig.Offiaccount) (provider *Provider, err error) {
	helperConf := &httphelper.Config{
		Config: &client.Config{
			Timeout: 30,
		},
		BaseUrl: endpoint,
	}
	h, err := httphelper.NewRequestHelper(helperConf)
	if err != nil {
		return nil, err
	}

	ts := newTokenSource(h, AuthConfig{
		AppID:          c.AppID,
		Secret:         c.Secret,
		UseStableToken: c.UseStableAccessToken,
	}, cachex.NewMemoryCache())

	return &Provider{
		Conf: c,
		H:    h,
		TS:   ts,
	}, nil
}

func (provider *Provider) SetCache(cache cachex.Cache) {
	provider.TS.cache = cache
}
