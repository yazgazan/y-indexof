
package start

import (
  "net/http"
)

type Cache struct{
  Objects   map[string]*Method
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

