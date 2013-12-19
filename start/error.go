
package start

import (
  "fmt"
)

type Error struct{
  What  string
  Code  int
}

func (e *Error) Error() string {
  return e.What
}

func MakeError(code int, msg string) *Error {
  return &Error{
    fmt.Sprintf("Error %d : %s", code, msg),
    code,
  }
}

func MakeError500(msg string) *Error {
  return MakeError(500, msg)
}

func MakeError404() *Error {
  return MakeError(404, "File not found")
}

