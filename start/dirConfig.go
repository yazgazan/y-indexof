
package start

import (
  "github.com/BurntSushi/toml"

  "io/ioutil"
)

type DirConfig struct{
  Method    string
  MethodId  int
  Forward   string
  View      string
}

func (c *DirConfig) MatchMethod() {
  if c.Method == "forward" {
    c.MethodId = Method_Forward
  } else if c.Method == "Index" {
    c.MethodId = Method_Index
  } else {
    c.MethodId = Method_Unknown
  }
}

type DirConfigError struct{
  What  string
}

func (e *DirConfigError) Error() string {
  return e.What
}

func (c *DirConfig) CheckDirConfig() error {
  if c.MethodId == Method_Forward {
    if c.Forward == "" {
      return &DirConfigError{
        "Method set to Forwarding but no port specified",
      }
    }
  }
  return nil
}

func MakeDirConfig(config Config) *DirConfig {
  return &DirConfig{
    MethodId: Method_Unknown,
    View: config.IndexView,
  }
}

func ReadDirConfig(path string, config Config) (*DirConfig, error) {
  conf := MakeDirConfig(config)

  content, err := ioutil.ReadFile(path)
  if err != nil {
    return nil, err
  }

  _, err = toml.Decode(string(content), &conf)
  if err != nil {
    return nil, err
  }
  conf.MatchMethod()

  err = conf.CheckDirConfig()
  if err != nil {
    return nil, err
  }

  return conf, nil
}

