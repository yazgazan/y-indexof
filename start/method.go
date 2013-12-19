
package start

import (
  "net/http"

  "path"
  "strings"
  "fmt"
  "os"
)

const (
  // Methodes types
  Method_Unknown =  iota
  Method_Static =   iota // serving file
  Method_Forward =  iota // forwarding request
  Method_Index =    iota // serving dir
  Method_Internal = iota // serving internal files (images, css, js ...)
)

type Method struct{
  Path      string // URL.Path
  MethodId  int // see Methodes types
  FullPath  string // e.g localhost:1243
  View      string // in case of Method_index
}

func (m *Method) FromDirConf(dirConf DirConfig, config Config) {
  m.MethodId = dirConf.MethodId
  m.View = dirConf.View
  if m.MethodId == Method_Forward {
    m.FullPath = dirConf.Forward
  } else {
    m.FullPath = path.Join(config.Root, m.Path)
  }
}

type ErrorFake struct{
  What string
}

func (e *ErrorFake) Error() string {
  return e.What
}

func GetFsPath(urlPath string, config Config) string {
  tmp_path := path.Join(config.Root, urlPath)
  return path.Clean(tmp_path)
}

func ConstructDirconfPath(config Config, urlPath string) string {
  tmp_path := path.Join(config.Root, urlPath)
  return path.Join(tmp_path, Config_file_name)
}

func ConstructDirConfig(urlPath string, config Config) *DirConfig {
  var conf *DirConfig
  dir, file := path.Split(urlPath)
  if file == "" {
    return MakeDirConfig(config)
  }
  if dir != "" && dir != "/" {
    conf = ConstructDirConfig(dir, config)
  } else {
    conf = MakeDirConfig(config)
  }

  newConf, err := ReadDirConfig(ConstructDirconfPath(config, urlPath), config)
  if err != nil {
    return conf
  }
  if conf.MethodId == Method_Forward {
    return conf
  }
  if newConf.MethodId != Method_Unknown {
    conf.Method = newConf.Method
    conf.MethodId = newConf.MethodId
  }
  if newConf.Forward != "" {
    conf.Forward = newConf.Forward
  }
  if newConf.View != "" {
    conf.View = newConf.View
  }
  return newConf
}

func MakeMethod() *Method {
  return &Method{
    "",
    Method_Unknown,
    "",
    "",
  }
}

func MatchInternal(urlPath string) bool {
  return strings.HasPrefix(urlPath, "/_/")
}

func GetMethod(
  w http.ResponseWriter,
  req *http.Request,
  config Config,
) (*Method, error) {
  var method = MakeMethod()
  method.Path = req.URL.Path

  match := MatchInternal(method.Path)
  if match == true {
    method.MethodId = Method_Internal
    method.FullPath = fmt.Sprintf("static/%s", method.Path[3:])
  } else {
    dirConf := ConstructDirConfig(req.URL.Path, config)

    method.FromDirConf(*dirConf, config)
  }

  if method.MethodId == Method_Forward {
    return method, nil
  }

  infos, err := os.Stat(method.FullPath)
  if err != nil {
    return nil, MakeError404()
  }

  if infos.Mode().IsDir() == true {
    method.MethodId = Method_Index
  } else if infos.Mode().IsRegular() == true {
    method.MethodId = Method_Static
  } else {
    return nil, MakeError500("unknown file mode")
  }

  return method, nil
}

