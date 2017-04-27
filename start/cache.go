/*
 * ----------------------------------------------------------------------------
 * "THE BEER-WARE LICENSE" (Revision 42):
 * <yazgazan@gmail.com> wrote this file. As long as you retain this notice you
 * can do whatever you want with this stuff. If we meet some day, and you think
 * this stuff is worth it, you can buy me a beer in return.
 * ----------------------------------------------------------------------------
 */

package start

import (
	"net/http"
	"sync"
)

type methodCache struct {
	Objects map[string]*methodConfig
	sync.RWMutex
}

func newCache() *methodCache {
	return &methodCache{
		Objects: make(map[string]*methodConfig),
	}
}

func getMethodFromCache(req *http.Request, cache *methodCache) *methodConfig {
	cache.RLock()
	method, ok := cache.Objects[req.URL.Path]
	cache.RUnlock()

	if ok {
		return method
	}

	return nil
}

func cacheSave(method *methodConfig, cache *methodCache) {
	cache.Lock()
	cache.Objects[method.Path] = method
	cache.Unlock()
}
