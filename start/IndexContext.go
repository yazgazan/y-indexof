
package start

import (
  "os"
  "encoding/json"
)

type IndexContext struct{
  Path            string
  FullPath        string
  DownloadPrefix  string
  Files           []IndexItem
  JsonContext     string
}

func (context *IndexContext) InitContext(method *Method, config Config) error {

  context.Path = method.Path
  context.FullPath = method.FullPath
  context.DownloadPrefix = config.DownloadPrefix

  if err := context.ReadDirInfos(config); err != nil {
    return err
  }

  context.CreateJson()

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

