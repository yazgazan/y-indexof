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
	"github.com/rakyll/magicmime"
	"net/http"
)

func HandleStatic(
	w http.ResponseWriter,
	req *http.Request,
	method *Method,
	config Config) error {
	err := magicmime.Open(magicmime.MAGIC_MIME_TYPE | magicmime.MAGIC_SYMLINK | magicmime.MAGIC_ERROR)
	if err != nil {
		return err
	}

	mimetype, err := magicmime.TypeByFile(method.FullPath)
	magicmime.Close()

	if err == nil && len(mimetype) != 0 {
		w.Header()["content-type"] = make([]string, 1)
		w.Header()["content-type"][0] = mimetype
	}
	method.ResolveType(config)
	if len(method.Type.Headers) != 0 {
		for key, values := range method.Type.Headers {
			w.Header()[key] = values
		}
	}
	http.ServeFile(w, req, method.FullPath)
	return nil
}
