
package start

import (
  "net/http"
)

func HandleStatic(w http.ResponseWriter, req *http.Request, method *Method) error {
  http.ServeFile(w, req, method.FullPath)
  return nil
}

func HandleMethod(w http.ResponseWriter, req *http.Request, method *Method) error {
  if method.MethodId == Method_Static {
    return HandleStatic(w, req, method)
  }
  if method.MethodId == Method_Internal {
    return HandleStatic(w, req, method)
  }
  if method.MethodId == Method_Forward {
    return MakeError(501, "fake error forward")
  }
  if method.MethodId == Method_Index {
    return MakeError(501, "fake error index")
  }
  return MakeError(501, "fake error")
}

