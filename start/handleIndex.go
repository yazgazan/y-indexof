
package start

import (
  "github.com/hoisie/mustache"

  "net/http"
  "fmt"
  "path"
)

func HandleIndex(
  w http.ResponseWriter,
  req *http.Request,
  method *Method,
  config Config) error {

  var context IndexContext

  context.InitSort(req)
  context.InitContext(method, config)
  view := path.Join(config.Views, method.View)
  res := mustache.RenderFile(view, context)
  fmt.Fprint(w, res)
  return nil
}

