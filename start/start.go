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

  "fmt"
)

func HandleError(w http.ResponseWriter, err error) {
  var code int

  msg := err.Error()
  customError, ok := err.(*Error)
  if ok == true {
    code = customError.Code
  } else {
    code = 500
  }
  http.Error(w, msg, code)
}

func Start(conf Config) error {
  var cache = MakeCache()

  fmt.Printf("%+v\n", conf)
  http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    var method *Method
    var err error

    method = GetMethodFromCache(req, cache)

    if method == nil || true {
      method, err = GetMethod(w, req, conf)
      if err != nil {
        HandleError(w, err)
        return
      }
      CacheSave(method, cache)
    } else {
      fmt.Println("from cache :D")
    }
    fmt.Printf("%+v\n", method)
    err = HandleMethod(w, req, method, conf)
    if err != nil {
      HandleError(w, err)
    }
  })
  err := http.ListenAndServe(conf.Listen, nil)
  return err
}

