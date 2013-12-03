
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
  initCmd.Flags().StringVarP(
    &initParams.File, "file", "f",
    M_init_file_default, M_init_file_help,
  )
  initCmd.Flags().StringVarP(
    &initParams.Url, "url", "u",
    M_init_url_default, M_init_url_help,
  )
  initCmd.Flags().StringVarP(
    &initParams.Dest, "dest", "d",
    M_init_dest_default, M_init_dest_help,
  )
  Cmd.AddCommand(initCmd)

  var startParams StartParams
  startCmd := &cobra.Command{
    Use: "start",
    Short: "Run the server",
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

