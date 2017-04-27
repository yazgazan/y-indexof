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

func handleError(w http.ResponseWriter, err error) {
	var code int

	msg := err.Error()
	if customError, ok := err.(*httpError); ok {
		code = customError.Code
	} else {
		code = 500
	}
	http.Error(w, msg, code)
}

// Start starts a server
func Start(conf Config) error {
	var cache = newCache()

	fmt.Printf("%+v\n", conf)
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		var method *methodConfig
		var err error

		method = getMethodFromCache(req, cache)

		if method == nil || true {
			method, err = getMethod(w, req, conf)
			if err != nil {
				handleError(w, err)
				return
			}
			cacheSave(method, cache)
		} else {
			fmt.Println("from cache :D")
		}
		fmt.Printf("%+v\n", method)
		err = handleMethod(w, req, method, conf)
		if err != nil {
			handleError(w, err)
		}
	})
	err := http.ListenAndServe(conf.Listen, nil)
	return err
}
