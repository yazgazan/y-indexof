

package start

import (
  "github.com/rakyll/magicmime"
  "net/http"
)

func HandleStatic(
  w http.ResponseWriter,
  req *http.Request,
  method *Method,
  config Config,) error {
  mimetype, err := magicmime.TypeByFile(method.FullPath)

  if err == nil && len(mimetype) != 0 {
    w.Header()["content-type"] = make([]string, 1)
    w.Header()["content-type"][0] = mimetype
  }
  method.ResolveType(config)
  if len(method.Type.Headers) != 0 {
    for key, values := range method.Type.Headers {
      w.Header()[key] = values
    }
  }
  http.ServeFile(w, req, method.FullPath)
  return nil
}

