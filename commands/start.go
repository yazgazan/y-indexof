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
	"github.com/yazgazan/y-indexof/start"

	"fmt"
	"os"
)

type StartParams struct {
	Listen string
	Dir    string
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
		fmt.Printf("Failed to start, couldn't load config : %s\n", err)
		return
	}

	// overriding config if needed
	if params.Listen != "" {
		conf.Listen = params.Listen
	}

	// start server
	start.Start(*conf)
}
