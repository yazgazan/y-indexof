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
    Use: "y-indexof",
    Short: M_yindexof_help,
  }

  Cmd.AddCommand(&cobra.Command{
    Use: "version",
    Short: M_version_help,
    Run: Version,
  })

  var initParams InitParams
  var initCmd = &cobra.Command{
    Use: "init",
    Short: M_init_help,
    Run: func (cmd *cobra.Command, args []string) {
      Init(cmd, args, initParams)
    },
  }
  initCmd.Flags().BoolVarP(
    &initParams.Local, "local", "l",
    M_init_local_default, M_init_local_help,
  )
  initCmd.Flags().BoolVarP(
    &initParams.Git, "git", "g",
    M_init_git_default, M_init_git_help,
  )
  initCmd.Flags().StringVarP(
    &initParams.File, "file", "f",
    M_init_file_default, M_init_file_help,
  )
  initCmd.Flags().StringVarP(
    &initParams.Url, "url", "u",
    M_init_url_default, M_init_url_help,
  )
  initCmd.Flags().StringVarP(
    &initParams.Branch, "branch", "b",
    M_init_branch_default, M_init_branch_help,
  )
  initCmd.Flags().StringVarP(
    &initParams.Dest, "dest", "d",
    M_init_dest_default, M_init_dest_help,
  )
  Cmd.AddCommand(initCmd)

  var startParams StartParams
  startCmd := &cobra.Command{
    Use: "start",
    Short: M_start_help,
    Run: func (cmd *cobra.Command, args []string) {
      Start(cmd, args, startParams)
    },
  }
  startCmd.Flags().StringVarP(
    &startParams.Listen, "listen", "l",
    M_start_listen_default, M_start_listen_help,
  )
  startCmd.Flags().StringVarP(
    &startParams.Dir, "dir", "d",
    M_start_dir_default, M_start_dir_help,
  )
  Cmd.AddCommand(startCmd)

  return Cmd
}

