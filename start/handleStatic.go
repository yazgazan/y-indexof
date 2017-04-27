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
	"gopkg.in/h2non/filetype.v0"
	"log"
	"net/http"
)

func HandleStatic(w http.ResponseWriter, req *http.Request, method *Method, config Config) error {
	var err error

	typ, err := filetype.MatchFile(method.FullPath)

	if err == nil && len(typ.MIME.Value) != 0 {
		log.Printf("mime-type detected for %q: %s\n", method.FullPath, typ.MIME.Value)
		w.Header()["content-type"] = make([]string, 1)
		w.Header()["content-type"][0] = typ.MIME.Value
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
