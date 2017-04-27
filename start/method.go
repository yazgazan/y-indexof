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
	"os"
	"path"
	"strings"
)

type methodType int

const (
	// Methodes types
	methodUnknown    methodType = iota
	methodStatic                // serving file
	methodIndex                 // serving dir
	methodInternal              // serving internal files (images, css, js ...)
	methodCustomView            // serving custom views
)

type methodConfig struct {
	Path     string     // URL.Path
	MethodID methodType // see Methodes types
	FullPath string     // e.g localhost:1243
	View     string     // in case of Method_index or Method_CustomView
	Type     fileType   // type config
}

func (m *methodConfig) FromDirConf(dirConf dirConfig, config Config) {
	m.View = path.Join(config.Views, dirConf.View)
	m.FullPath = path.Join(config.Root, m.Path)
}

func (m *methodConfig) ResolveType(config Config) {
	allType, ok := config.Types["All"]
	if ok {
		m.Type = allType
	}
	if m.MethodID == methodIndex {
		_type, ok := config.Types["Folder"]
		if ok {
			m.Type.Merge(_type)
		}
	} else {
		defaultType, ok := config.Types["Default"]
		if ok {
			m.Type.Merge(defaultType)
		}
		var ext string
		if m.MethodID == methodCustomView {
			ext = path.Ext(m.View)
		} else {
			ext = path.Ext(m.Path)
		}
		if ext == "" {
			return
		}
		for _, curType := range config.Types {
			for _, typeExt := range curType.Exts {
				if typeExt == ext {
					m.Type.Merge(curType)
					return
				}
			}
		}
	}
}

func constructDirconfPath(config Config, urlPath string) string {
	return path.Join(path.Join(config.Root, urlPath), ConfigFileName)
}

func constructDirConfig(urlPath string, config Config) *dirConfig {
	var conf *dirConfig

	dir, file := path.Split(urlPath)
	if dir != "/" && file != "" {
		conf = constructDirConfig(dir, config)
	} else {
		conf = newDirConfig(config)
	}

	newConf, err := readDirConfig(constructDirconfPath(config, urlPath), config)
	if err != nil {
		return conf
	}
	if newConf.View != "" {
		conf.View = newConf.View
	}
	return newConf
}

func newMethodConfig() *methodConfig {
	return &methodConfig{
		MethodID: methodUnknown,
	}
}

func matchCustomView(urlPath string, config Config) bool {
	_, name := path.Split(urlPath)

	if len(name) == 0 {
		return false
	}
	for key := range config.CustomViews {
		if key == name {
			return true
		}
	}
	return false
}

func customViewExtractPath(urlPath string, config Config) string {
	_, name := path.Split(urlPath)

	if len(name) == 0 {
		return ""
	}
	for key, val := range config.CustomViews {
		if key == name {
			return path.Join(config.Views, path.Join("custom/", val))
		}
	}
	return ""
}

func matchInternal(urlPath string) bool {
	return strings.HasPrefix(urlPath, "/_/")
}

func getMethod(w http.ResponseWriter, req *http.Request, config Config) (*methodConfig, error) {
	if matchInternal(req.URL.Path) {
		return getInternalMethod(w, req, config)
	}

	return getExternalMethod(w, req, config)
}

func getExternalMethod(w http.ResponseWriter, req *http.Request, config Config) (*methodConfig, error) {
	method := newMethodConfig()
	method.Path = req.URL.Path

	isCustomView := matchCustomView(method.Path, config)

	dirConf := constructDirConfig(req.URL.Path, config)
	method.FromDirConf(*dirConf, config)

	if isCustomView {
		method.View = customViewExtractPath(method.Path, config)
		method.FullPath, _ = path.Split(method.FullPath)
	}

	infos, err := os.Stat(method.FullPath)
	if err != nil {
		return nil, notFoundError()
	}

	if isCustomView && infos.Mode().IsDir() {
		method.MethodID = methodCustomView
	} else if infos.Mode().IsDir() {
		method.MethodID = methodIndex
	} else if infos.Mode().IsRegular() {
		method.MethodID = methodStatic
	} else if isCustomView {
		return nil, internalError("custom view not found")
	} else {
		return nil, internalError("unknown file mode")
	}

	method.ResolveType(config)

	return method, nil
}

func getInternalMethod(w http.ResponseWriter, req *http.Request, config Config) (*methodConfig, error) {
	var isCustomView = false
	var method = newMethodConfig()
	method.Path = req.URL.Path

	method.MethodID = methodInternal
	method.FullPath = fmt.Sprintf("static/%s", method.Path[3:])

	infos, err := os.Stat(method.FullPath)
	if err != nil {
		return nil, notFoundError()
	}

	if isCustomView && infos.Mode().IsDir() {
		method.MethodID = methodCustomView
	} else if infos.Mode().IsRegular() {
		method.MethodID = methodStatic
	}

	return nil, accessForbiddenError()
}
