/*
 * ----------------------------------------------------------------------------
 * "THE BEER-WARE LICENSE" (Revision 42):
 * <yazgazan@gmail.com> wrote this file. As long as you retain this notice you
 * can do whatever you want with this stuff. If we meet some day, and you think
 * this stuff is worth it, you can buy me a beer in return.
 * ----------------------------------------------------------------------------
 */

package commands

import (
	"github.com/spf13/cobra"
	initlib "github.com/yazgazan/y-indexof/init"

	"fmt"
)

type initParams struct {
	Local  bool
	Git    bool
	File   string
	URL    string
	Branch string
	Dest   string
}

func runInit(cmd *cobra.Command, args []string, params initParams) {
	var err error

	if params.Git && !params.Local {
		err = initlib.GitInit(params.URL, params.Dest, params.Branch)
		if err != nil {
			fmt.Println("Error fetching git repo.")
		}
		return
	}
	if !params.Local {
		err = initlib.DownloadInit(params.URL, params.File)
		if err != nil {
			fmt.Println("Error downloading init.tar")
			return
		}
	}
	err = initlib.Extract(params.File, params.Dest)
	if err != nil {
		fmt.Println("Error extracting init.tar")
	}
}
