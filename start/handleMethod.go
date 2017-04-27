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

func handleMethod(w http.ResponseWriter, req *http.Request, method *methodConfig, config Config) error {
	if method.MethodID == methodStatic {
		return handleStatic(w, req, method, config)
	}
	if method.MethodID == methodInternal {
		return handleStatic(w, req, method, config)
	}
	if method.MethodID == methodIndex {
		return handleIndex(w, req, method, config)
	}
	if method.MethodID == methodCustomView {
		return handleIndex(w, req, method, config)
	}
	return genericError(500, "Method not found (should not happen)")
}
