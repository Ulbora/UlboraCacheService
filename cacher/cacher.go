package cacher

import "github.com/coocood/freecache"

//Cacher the cache
type Cacher struct {
	cache *freecache.Cache
}

//NewCache initial cache
func (c *Cacher) NewCache(cacheSize int) bool {
	var rtn = false
	c.cache = freecache.NewCache(cacheSize)
	if c.cache != nil {
		rtn = true
	}
	return rtn
}

//Set set a cache item
func (c *Cacher) Set(key, val string, expireSeconds int) error {
	return c.cache.Set([]byte(key), []byte(val), expireSeconds)
}

//Get get cache item
func (c *Cacher) Get(key string) (string, error) {
	rtn, err := c.cache.Get([]byte(key))
	return string(rtn), err
}

//SetObj set a cache item as object
func (c *Cacher) SetObj(key string, val []byte, expireSeconds int) error {
	return c.cache.Set([]byte(key), val, expireSeconds)
}

//GetObj get cache item as object
func (c *Cacher) GetObj(key string) ([]byte, error) {
	return c.cache.Get([]byte(key))
}

//Delete remove cache item
func (c *Cacher) Delete(key string) bool {
	return c.cache.Del([]byte(key))
}
