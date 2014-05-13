
package start

import (
  "github.com/hoisie/mustache"

  "net/http"
  "fmt"
)

func HandleIndex(
  w http.ResponseWriter,
  req *http.Request,
  method *Method,
  config Config) error {

  var context IndexContext

  context.InitSort(req)
  context.InitContext(method, config)
  view := method.View
  res := mustache.RenderFile(view, context)
  fmt.Fprint(w, res)
  return nil
}

