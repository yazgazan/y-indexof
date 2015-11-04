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
)

type Cache struct {
	Objects map[string]*Method
}

func MakeCache() *Cache {
	return &Cache{
		make(map[string]*Method),
	}
}

func GetMethodFromCache(req *http.Request, cache *Cache) *Method {
	method, ok := cache.Objects[req.URL.Path]

	if ok == false {
		return nil
	}

	return method
}

func CacheSave(method *Method, cache *Cache) {
	cache.Objects[method.Path] = method
}
