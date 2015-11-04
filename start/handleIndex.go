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
	"github.com/hoisie/mustache"

	"fmt"
	"net/http"
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
