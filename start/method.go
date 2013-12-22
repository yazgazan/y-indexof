
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
  Method_Index =    iota // serving dir
  Method_Internal = iota // serving internal files (images, css, js ...)
)

const (
  // Responses types
  Response_html = iota // use the view to render html index
  Response_json = iota // respond a json listing - TODO Later
  // more to come ...
)

type Method struct{
  Path      string // URL.Path
  MethodId  int // see Methodes types
  FullPath  string // e.g localhost:1243
  View      string // in case of Method_index
  ResType   int // see Response types
  Type      Type // type config
}

func (m *Method) FromDirConf(dirConf DirConfig, config Config) {
  m.View = dirConf.View
  m.FullPath = path.Join(config.Root, m.Path)
}

func (m *Method) ResolveType(config Config) {
  allType, ok := config.Types["All"]
  if ok == true {
    m.Type = allType
  }
  if m.MethodId == Method_Index {
    _type, ok := config.Types["Folder"]
    if ok == true {
      m.Type.Merge(_type)
    }
  } else {
    defaultType, ok := config.Types["Default"]
    if ok == true {
      m.Type.Merge(defaultType)
    }
    ext := path.Ext(m.Path)
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
  if dir != "/" && file != ""{
    conf = ConstructDirConfig(dir, config)
  } else {
    conf = MakeDirConfig(config)
  }

  newConf, err := ReadDirConfig(ConstructDirconfPath(config, urlPath), config)
  if err != nil {
    return conf
  }
  if newConf.View != "" {
    conf.View = newConf.View
  }
  return newConf
}

func MakeMethod() *Method {
  return &Method{
    MethodId: Method_Unknown,
    ResType: Response_html,
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

  infos, err := os.Stat(method.FullPath)
  if err != nil {
    return nil, MakeError404()
  }

  if infos.Mode().IsDir() == true && match == false {
    method.MethodId = Method_Index
  } else if infos.Mode().IsRegular() == true {
    method.MethodId = Method_Static
  } else if match == true {
    return nil, MakeError403()
  } else {
    return nil, MakeError500("unknown file mode")
  }

  method.ResolveType(config)

  return method, nil
}

