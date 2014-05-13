
package start

import (
  "net/http"
)

func HandleMethod(
  w http.ResponseWriter,
  req *http.Request,
  method *Method,
  config Config,) error {
  if method.MethodId == Method_Static {
    return HandleStatic(w, req, method, config)
  }
  if method.MethodId == Method_Internal {
    return HandleStatic(w, req, method, config)
  }
  if method.MethodId == Method_Index {
    return HandleIndex(w, req, method, config)
  }
  if method.MethodId == Method_CustomView {
    return HandleIndex(w, req, method, config)
  }
  return MakeError(500, "Method not found (should not happen)")
}

