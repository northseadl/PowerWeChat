package weconfig

import "fmt"

type Redis struct {
	Addr     string // redis的地址
	Password string // redis的密码
	DB       int    // redis的数据库
}

func (c *Redis) Default() {
}

func (c *Redis) Validate() error {
	if c.Addr == "" {
		return fmt.Errorf("miss addr")
	}
	return nil
}
