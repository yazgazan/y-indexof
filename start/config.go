
package start

import (
  "github.com/BurntSushi/toml"

  "io/ioutil"
)

type TypeHeader struct{
  Key     string
  Value   string
  Values  []string
}

type Type struct{
  Exts        []string
  Image       string
  Headers     map[string][]string
}

func (base *Type) Merge(extend Type) {
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

type Config struct{
  Listen          string
  Root            string
  Views           string
  IndexView       string
  Static          string
  DownloadPrefix  string
  Types           map[string]Type
  UserDefined     map[string]string
}

func MakeConfig() *Config {
  return &Config{
    Listen: default_listen,
    Root: default_root,
    Views: default_views,
    IndexView: default_indexView,
    Static: default_static,
    DownloadPrefix: default_download_prefix,
  }
}

func ReadConfig(path string) (*Config, error) {
  conf := MakeConfig()

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

