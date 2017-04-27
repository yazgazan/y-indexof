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
)

func setupCmd() *cobra.Command {
	var Cmd = &cobra.Command{
		Use:   "y-indexof",
		Short: yindexofHelp,
	}

	Cmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: versionHelp,
		Run:   version,
	})

	var initP initParams
	var initCmd = &cobra.Command{
		Use:   "init",
		Short: initHelp,
		Run: func(cmd *cobra.Command, args []string) {
			runInit(cmd, args, initP)
		},
	}
	initCmd.Flags().BoolVarP(
		&initP.Local, "local", "l",
		initLocalDefault, initLocalHelp,
	)
	initCmd.Flags().BoolVarP(
		&initP.Git, "git", "g",
		initGitDefault, initGitHelp,
	)
	initCmd.Flags().StringVarP(
		&initP.File, "file", "f",
		initFileDefault, initFileHelp,
	)
	initCmd.Flags().StringVarP(
		&initP.URL, "url", "u",
		initURLDefault, initURLHelp,
	)
	initCmd.Flags().StringVarP(
		&initP.Branch, "branch", "b",
		initBranchDefault, initBranchHelp,
	)
	initCmd.Flags().StringVarP(
		&initP.Dest, "dest", "d",
		initDestDefault, initDestHelp,
	)
	Cmd.AddCommand(initCmd)

	var startP startParams
	startCmd := &cobra.Command{
		Use:   "start",
		Short: startHelp,
		Run: func(cmd *cobra.Command, args []string) {
			runStart(cmd, args, startP)
		},
	}
	startCmd.Flags().StringVarP(
		&startP.Listen, "listen", "l",
		startListenDefault, startListenHelp,
	)
	startCmd.Flags().StringVarP(
		&startP.Dir, "dir", "d",
		startDirDefault, startDirHelp,
	)
	Cmd.AddCommand(startCmd)

	return Cmd
}
