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

type httpError struct {
	What string
	Code int
}

func (e *httpError) Error() string {
	return e.What
}

func genericError(code int, msg string) *httpError {
	return &httpError{
		fmt.Sprintf("Error %d : %s", code, msg),
		code,
	}
}

func internalError(msg string) *httpError {
	return genericError(500, msg)
}

func notFoundError() *httpError {
	return genericError(404, "File not found")
}

func accessForbiddenError() *httpError {
	return genericError(403, "Access Forbidden")
}
