
package commands

import (
  "github.com/yazgazan/y-indexof/start"
  "github.com/spf13/cobra"

  "fmt"
  "os"
)

type StartParams struct{
  Listen        string
  Dir           string
}

func Start(cmd *cobra.Command, args []string, params StartParams) {

  // move to Dir
  err := os.Chdir(params.Dir)
  if err != nil {
    fmt.Println("Failed to start, couldn't cd into", params.Dir)
    return
  }

  // read config
  conf, err := start.ReadConfig(start.Config_file_name)
  if err != nil {
    fmt.Println("Failed to start, couldn't load config")
    return
  }

  // overriding config if needed
  if params.Listen != "" {
    conf.Listen = params.Listen
  }

  // start server
  start.Start(*conf)
}

