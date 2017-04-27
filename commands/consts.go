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
	yindexofHelp    = "y-indexof is a simple yet modular indexof server"
	versionHelp     = "Print programm version"
	initHelp        = "Init the layout in the current directory"
	initLocalHelp   = "Use a local init.tar"
	initGitHelp     = "Use a git as init folder"
	initFileHelp    = "Path to init.tar (for local and download)"
	initURLHelp     = "Where to download init.tar from"
	initBranchHelp  = "Which branch to checkout (git only)"
	initDestHelp    = "Where to extract init.tar"
	startHelp       = "Run the server"
	startListenHelp = "Port to listen (override config port)"
	startDirHelp    = "Directory containing the configurations and files"

	initLocalDefault   = false
	initGitDefault     = true
	initFileDefault    = "/tmp/init.tar"
	initURLDefault     = "git@github.com:yazgazan/y-indexof_init.git"
	initBranchDefault  = "master"
	initDestDefault    = "."
	startListenDefault = ""
	startDirDefault    = "."
)
