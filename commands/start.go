
package commands

import (
  /* "github.com/yazgazan/y-indexof/start" */
  "github.com/spf13/cobra"

  "fmt"
)

type StartParams struct{
  Listen        string
  Dir           string
}

func Start(cmd *cobra.Command, args []string, params StartParams) {
  fmt.Println("Starting ...");
  fmt.Println(params.Listen)
  fmt.Println(params.Dir)
}

