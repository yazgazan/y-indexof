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
  "os"
  "sort"
  "encoding/json"
)

type IndexContext struct{
  Path            string
  FullPath        string
  DownloadPrefix  string
  Files           []IndexItem
  JsonContext     string
  Sort            string
  UserDefined     map[string]string
}

func (context *IndexContext) InitSort(req *http.Request) {
  cookie, err := req.Cookie("sort")
  if err != nil {
    context.Sort = "alpha"
    return
  }
  switch cookie.Value {
    case "date": context.Sort = "date"
    case "size": context.Sort = "size"
    case "alpha": context.Sort = "alpha"
  }
  if context.Sort == "" {
    context.Sort = "alpha"
  }
}

func (context *IndexContext) InitContext(method *Method, config Config) error {
  fullPath := method.FullPath

  context.Path = method.Path
  context.FullPath = method.FullPath
  context.DownloadPrefix = config.DownloadPrefix
  context.UserDefined = config.UserDefined

  if err := context.ReadDirInfos(config); err != nil {
    return err
  }

  switch context.Sort {
    case "date": sort.Sort(DateSortFiles(context.Files))
    case "alpha": sort.Sort(AlphaSortFiles(context.Files))
    case "size": sort.Sort(SizeSortFiles(context.Files))
  }

  if config.ShowFullPath == false {
    context.FullPath = context.Path
    for id := range context.Files {
      context.Files[id].FullPath = context.Files[id].Path
    }
  }
  context.CreateJson()
  if config.ShowFullPath == false {
    context.FullPath = fullPath
    for id := range context.Files {
      context.Files[id].FullPath = context.Files[id].fullPath
    }
  }

  return nil
}

func (context *IndexContext) CreateJson() {
  bytes, err := json.Marshal(context)

  if err != nil {
    panic(err) // shouldn't happen
  }
  context.JsonContext = string(bytes)
}

func IsExcluded(file os.FileInfo) bool {
  name := file.Name()

  if name[0] == '.' {
    return true
  }
  if name == "indexof.toml" {
    return true
  }

  return false
}

func CountFiles(files []os.FileInfo) int {
  count := 0

  for _, info := range files {
    if IsExcluded(info) == true {
      continue
    }
    count += 1
  }

  return count
}

func (context *IndexContext) ReadDirInfos(config Config) error {
  file, err := os.Open(context.FullPath)

  if err != nil {
    return err
  }

  fileInfos, err := file.Readdir(0)
  if err != nil {
    return err
  }
  filesCount := CountFiles(fileInfos)
  context.Files = make([]IndexItem, filesCount)

  i := 0
  for _, info := range fileInfos {
    if IsExcluded(info) == true {
      continue
    }
    context.Files[i].Populate(info, context, config)
    i++
  }
  return nil
}

