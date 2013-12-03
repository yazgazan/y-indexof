
package init

import (
  "archive/tar"
  "os"
  "io"
  "fmt"
  "path"
)

func extractDir(path string, mode os.FileMode) error {
  err := os.MkdirAll(path, mode)
  return err
}

func extractFile(path string, mode os.FileMode, fileIn io.Reader) error {
  file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
  if err != nil {
    return err
  }
  _, err = io.Copy(file, fileIn)
  if err != nil {
    return err
  }
  return nil
}

func Extract(filename string, dest string) error {
  file, err := os.Open(filename)
  if err != nil {
    return err
  }
  tr := tar.NewReader(file)
  for {
    header, err := tr.Next()
    if err == io.EOF {
      break
    }
    if err != nil {
      return err
    }

    fullPath := path.Join(dest, header.Name)
    mode := header.FileInfo().Mode()

    fmt.Println(header.Name)
    if mode.IsDir() {
      fmt.Printf("Extracting %s ...\n", header.Name)
      err = extractDir(fullPath, mode)
      if err != nil {
        return err
      }
    } else if mode.IsRegular() {
      fmt.Printf("Extracting %s ...\n", header.Name)
      err = extractFile(fullPath, mode, tr)
      if err != nil {
        return err
      }
    } else {
      fmt.Printf("Don't know how to extract %s.\n", header.Name)
    }
  }
  err = file.Close()
  if err != nil {
    panic(err)
  }
  return nil
}

