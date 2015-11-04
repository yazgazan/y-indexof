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
	"fmt"
)

type Error struct {
	What string
	Code int
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

func MakeError403() *Error {
	return MakeError(404, "Access Forbiden")
}
