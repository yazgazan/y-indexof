
package init

import (
  "net/http"
  "fmt"
  "io"
  "os"
)

func DownloadInit(url string, outpath string) error {
  fmt.Println("ha !")
  outfile, err := os.Create(outpath)
  if err != nil {
    return err
  }

  res, err := http.Get(url)
  if err != nil {
    return err
  }

  _, err = io.Copy(outfile, res.Body)
  if err != nil {
    return err
  }
  outfile.Close()

  return nil
}

