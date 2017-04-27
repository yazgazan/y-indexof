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

	"fmt"
)

func version(cmd *cobra.Command, args []string) {
	fmt.Println("y-indexof v0.0.1")
}
