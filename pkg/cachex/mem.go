package cachex

import (
	"errors"
	"sync"
	"time"
)

type item struct {
	value      string
	expiration int64
}

type MemoryCache struct {
	cache map[string]*item
	mu    sync.RWMutex
}

func NewMemoryCache() Cache {
	m := &MemoryCache{
		cache: make(map[string]*item),
	}
	go m.ttlScanner()
	return m
}

func (m *MemoryCache) Set(key string, value string, ttl time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache[key] = &item{
		value:      value,
		expiration: time.Now().Add(ttl).UnixNano(),
	}
	return nil
}

func (m *MemoryCache) Get(key string) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if item, found := m.cache[key]; found {
		if time.Now().UnixNano() > item.expiration {
			delete(m.cache, key)
			return "", errors.New("cache expired")
		}
		return item.value, nil
	}
	return "", errors.New("cache not found")
}

func (m *MemoryCache) ttlScanner() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		m.mu.Lock()
		for key, item := range m.cache {
			if time.Now().UnixNano() > item.expiration {
				delete(m.cache, key)
			}
		}
		m.mu.Unlock()
	}
}
