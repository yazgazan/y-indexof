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
	"github.com/BurntSushi/toml"

	"io/ioutil"
)

type fileType struct {
	Exts    []string
	Image   string
	Headers map[string][]string
}

func (base *fileType) Merge(extend fileType) {
	if len(extend.Image) != 0 {
		base.Image = extend.Image
	}
	if len(extend.Exts) != 0 {
		base.Exts = extend.Exts
	}
	if len(extend.Headers) != 0 {
		if len(base.Headers) == 0 {
			base.Headers = make(map[string][]string)
		}
		for key, values := range extend.Headers {
			base.Headers[key] = values
		}
	}
}

type Config struct {
	Listen         string
	Root           string
	Views          string
	IndexView      string
	Static         string
	DownloadPrefix string
	ShowFullPath   bool
	Types          map[string]fileType
	UserDefined    map[string]string
	CustomViews    map[string]string
}

func newConfig() *Config {
	return &Config{
		Listen:         defaultListen,
		Root:           defaultRoot,
		Views:          defaultViews,
		IndexView:      defaultIndexView,
		Static:         defaultStatic,
		DownloadPrefix: defaultDownloadPrefix,
		ShowFullPath:   false,
	}
}

func ReadConfig(path string) (*Config, error) {
	conf := newConfig()

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	_, err = toml.Decode(string(content), &conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
