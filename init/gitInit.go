/*
 * ----------------------------------------------------------------------------
 * "THE BEER-WARE LICENSE" (Revision 42):
 * <yazgazan@gmail.com> wrote this file. As long as you retain this notice you
 * can do whatever you want with this stuff. If we meet some day, and you think
 * this stuff is worth it, you can buy me a beer in return.
 * ----------------------------------------------------------------------------
 */

package init

import (
  "github.com/yazgazan/go-vcs"

  "fmt"
)

func GitInit(url string, dest string, branch string) error {
  fmt.Println("Cloning ", url, " to ", dest, " ...")
  repo, err := vcs.Git.Clone(url, dest)

  if err != nil {
    return err
  }

  fmt.Println("Checkout ", branch, " ...")
  _, err = repo.CheckOut(branch)

  return err
}

