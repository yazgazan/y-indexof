
package commands

import (
  initlib "github.com/yazgazan/y-indexof/init"
  "github.com/spf13/cobra"

  "fmt"
)

type InitParams struct{
  Local bool
  File  string
  Url   string
  Dest  string
}

func Init(cmd *cobra.Command, args []string, params InitParams) {
  var err error

  if params.Local == false {
    err = initlib.DownloadInit(params.Url, params.File)
    if err != nil {
      fmt.Println("Error downloading init.tar")
      return
    }
  }
  err = initlib.Extract(params.File, params.Dest)
  if err != nil {
    fmt.Println("Error extracting init.tar")
    return
  }
}

