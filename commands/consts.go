/*
 * ----------------------------------------------------------------------------
 * "THE BEER-WARE LICENSE" (Revision 42):
 * <yazgazan@gmail.com> wrote this file. As long as you retain this notice you
 * can do whatever you want with this stuff. If we meet some day, and you think
 * this stuff is worth it, you can buy me a beer in return.
 * ----------------------------------------------------------------------------
 */

package commands

const (
  M_yindexof_help = "y-indexof is a simple yet modular indexof server"
  M_version_help = "Print programm version"
  M_init_help = "Init the layout in the current directory"
  M_init_local_help = "Use a local init.tar"
  M_init_git_help = "Use a git as init folder"
  M_init_file_help = "Path to init.tar (for local and download)"
  M_init_url_help = "Where to download init.tar from"
  M_init_branch_help = "Which branch to checkout (git only)"
  M_init_dest_help = "Where to extract init.tar"
  M_start_help = "Run the server"
  M_start_listen_help = "Port to listen (override config port)"
  M_start_dir_help = "Directory containing the configurations and files"

  M_init_local_default = false
  M_init_git_default = true
  M_init_file_default = "/tmp/init.tar"
  M_init_url_default = "git@github.com:yazgazan/y-indexof_init.git"
  M_init_branch_default = "master"
  M_init_dest_default = "."
  M_start_listen_default = ""
  M_start_dir_default = "."
)

