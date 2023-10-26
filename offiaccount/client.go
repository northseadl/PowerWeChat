package offiaccount

import (
	"fmt"

	"github.com/northseadl/PowerWeChat/pro/offiaccount/core"
	"github.com/northseadl/PowerWeChat/pro/pkg/cachex"
	"github.com/northseadl/PowerWeChat/pro/weconfig"
)

type API struct {
	p *core.Provider
}

type Option func(*API)

func WithCache(cache cachex.Cache) Option {
	return func(api *API) {
		api.p.SetCache(cache)
	}
}

func New(c weconfig.Offiaccount, opts ...Option) (api *API, err error) {
	provider, err := core.NewProvider(&c)
	if err != nil {
		return nil, fmt.Errorf("init api failed: %w", err)
	}

	for _, opt := range opts {
		opt(&API{provider})
	}
	return &API{provider}, nil
}

func (api *API) GetToken() (string, error) {
	return api.p.TS.GetAccessToken()
}
