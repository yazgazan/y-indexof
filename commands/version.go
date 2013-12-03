
package commands

import (
  "github.com/spf13/cobra"

  "fmt"
)

func Version(cmd *cobra.Command, args []string) {
  fmt.Println("y-indexof v0.0.1");
}

