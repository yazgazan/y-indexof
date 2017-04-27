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
	"io"
	"net/http"

	"github.com/hoisie/mustache"
)

func handleIndex(w io.Writer, req *http.Request, method *methodConfig, config Config) error {
	var context indexContext

	context.InitSort(req)
	context.InitContext(method, config)
	view := method.View
	res := mustache.RenderFile(view, context)
	_, err := fmt.Fprint(w, res)
	return err
}
