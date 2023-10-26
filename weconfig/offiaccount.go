package weconfig

import "fmt"

type Offiaccount struct {
	AppID                string // 微信开放平台的AppID
	Secret               string // 微信开放平台的Secret
	Token                string // 微信开放平台的Token
	AESKey               string // 微信开放平台的AESKey
	Callback             string // 微信开放平台的回调地址
	UseStableAccessToken bool   // 是否使用稳定的access_token
}

// Config 是微信开放平台的配置信息

// SetDefaultValues 设置默认值
func (c *Offiaccount) Default() {

}

// Validate 校验配置信息是否有效
func (c *Offiaccount) Validate() error {
	if c.AppID == "" {
		return fmt.Errorf("miss appId")
	}
	if c.Secret == "" {
		return fmt.Errorf("miss secret")
	}
	return nil
}
