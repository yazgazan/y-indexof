
package init

import (
  "github.com/sourcegraph/go-vcs"

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

