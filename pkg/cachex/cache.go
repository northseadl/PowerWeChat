package cachex

import (
	"time"
)

type Cache interface {
	Set(key string, value string, ttl time.Duration) error
	Get(key string) (string, error)
}
