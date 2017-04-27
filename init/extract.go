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
	"archive/tar"
	"fmt"
	"io"
	"os"
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

	return err
}

func Extract(filename string, dst string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	tr := tar.NewReader(file)
	return extractTar(tr, dst)
}

func extractTar(tr *tar.Reader, dst string) error {
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		fullPath := path.Join(dst, header.Name)
		mode := header.FileInfo().Mode()

		fmt.Println("Extracting", header.Name, "...")
		if mode.IsDir() {
			err = extractDir(fullPath, mode)
			if err != nil {
				return err
			}
		} else if mode.IsRegular() {
			err = extractFile(fullPath, mode, tr)
			if err != nil {
				return err
			}
		} else if header.Typeflag == tar.TypeLink {
			continue // TODO : see what it is for ...
		} else {
			fmt.Println("Don't know how to extract", header.Name)
			fmt.Println("File Type :", header.Typeflag)
			fmt.Println("Report the error (with file type) to yazgazan@gmail.com")
		}
	}

	return nil
}
