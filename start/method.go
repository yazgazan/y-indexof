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
  Method_CustomView = iota // serving custom views
)

type Method struct{
  Path      string // URL.Path
  MethodId  int // see Methodes types
  FullPath  string // e.g localhost:1243
  View      string // in case of Method_index or Method_CustomView
  Type      Type // type config
}

func (m *Method) FromDirConf(dirConf DirConfig, config Config) {
  m.View = path.Join(config.Views, dirConf.View)
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
    var ext string
    if m.MethodId == Method_CustomView {
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
  }
}

func MatchCustomView(urlPath string, config Config) bool {
  _, name := path.Split(urlPath)

  if len(name) == 0 {
    return false
  }
  for key, _ := range config.CustomViews {
    if key == name {
      return true
    }
  }
  return false
}

func CustomViewExtractPath(urlPath string, config Config) string {
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

func MatchInternal(urlPath string) bool {
  return strings.HasPrefix(urlPath, "/_/")
}

func GetMethod(
  w http.ResponseWriter,
  req *http.Request,
  config Config,
) (*Method, error) {
  var isCustomView = false
  var method = MakeMethod()
  method.Path = req.URL.Path

  match := MatchInternal(method.Path)
  if match == true {
    method.MethodId = Method_Internal
    method.FullPath = fmt.Sprintf("static/%s", method.Path[3:])
  } else {
    isCustomView = MatchCustomView(method.Path, config)

    dirConf := ConstructDirConfig(req.URL.Path, config)
    method.FromDirConf(*dirConf, config)

    if isCustomView == true {
      method.View =  CustomViewExtractPath(method.Path, config)
      method.FullPath, _ = path.Split(method.FullPath)
    }
  }

  infos, err := os.Stat(method.FullPath)
  if err != nil {
    return nil, MakeError404()
  }

  if isCustomView == true && infos.Mode().IsDir() == true {
    method.MethodId = Method_CustomView
  } else if infos.Mode().IsDir() == true && match == false {
    method.MethodId = Method_Index
  } else if infos.Mode().IsRegular() == true {
    method.MethodId = Method_Static
  } else if match == true {
    return nil, MakeError403()
  } else if isCustomView == true {
    return nil, MakeError500("custom view not found")
  } else {
    return nil, MakeError500("unknown file mode")
  }

  method.ResolveType(config)

  return method, nil
}

