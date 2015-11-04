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
)

func HandleMethod(
	w http.ResponseWriter,
	req *http.Request,
	method *Method,
	config Config) error {
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
