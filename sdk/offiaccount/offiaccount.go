package offiaccount

import (
	"github.com/artisancloud/httphelper"
)

type Client struct {
	config *OfficialAccountConfig
	helper httphelper.Helper
}

func NewClient(config *OfficialAccountConfig) (client *Client, err error) {
	helper, err := httphelper.NewRequestHelper(&httphelper.Config{
		BaseUrl: "https://api.weixin.qq.com/",
	})
	if err != nil {
		panic(err)
	}
	client = &Client{
		config: config,
		helper: helper,
	}
	return client, nil
}
