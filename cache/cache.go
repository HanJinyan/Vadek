package cache

import (
	goCache "github.com/patrickmn/go-cache"
	"time"
)

type Cache interface {
	SetDafault(key string, value interface{})
	//expiration time.Duration 过期时间，持续时间
	Set(key string, value interface{}, expiration time.Duration)
	Get(key string) (interface{}, bool)
	Delete(key string)
	//批量删除
	BatchDelete(keys []string)
}

// TODO 修改vadekCache
type cacheTemp struct {
	vadekCache *goCache.Cache
}

func (c *cacheTemp) SetDafault(key string, value interface{}) {
	c.vadekCache.SetDefault(key, value)
}

func (c *cacheTemp) Set(key string, value interface{}, expiration time.Duration) {
	c.vadekCache.Set(key, value, expiration)
}

func (c *cacheTemp) Get(key string) (interface{}, bool) {
	return c.vadekCache.Get(key)
}
func (c *cacheTemp) Delete(key string) {
	c.vadekCache.Delete(key)

}
func (c *cacheTemp) BatchDelete(keys []string) {
	for _, key := range keys {
		c.Delete(key)
	}

}

var _Cache = &cacheTemp{}

func NewCache() Cache {
	return &cacheTemp{
		vadekCache: goCache.New(time.Hour, time.Hour),
	}

}
