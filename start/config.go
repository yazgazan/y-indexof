
package start

import (
  "github.com/BurntSushi/toml"

  "io/ioutil"
)

type Config struct{
  Listen      string
  Root        string
  Views       string
  IndexView   string
  Static      string
}

func MakeConfig() *Config {
  return &Config{
    default_listen,
    default_root,
    default_views,
    default_indexView,
    default_static,
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

