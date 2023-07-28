package offiaccount

import (
	"github.com/pkg/errors"
)

type OfficialAccountConfig struct {
	AppID     string
	Secret    string
	AESKey    string
	HttpDebug bool
}

func (c *OfficialAccountConfig) Default() {
	// not to do anything
}

func (c *OfficialAccountConfig) Validate() error {
	if c.AppID == "" {
		return errors.New("AppID 不能为空")
	}
	if c.Secret == "" || c.Secret == "[app secret]" {
		return errors.New("Secret 不能为空")
	}
	return nil
}
