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
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadInit(url string, outpath string) error {
	outfile, err := os.Create(outpath)
	if err != nil {
		return err
	}
	defer outfile.Close()

	fmt.Println("Downloading tarball from ", url, "...")
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	_, err = io.Copy(outfile, res.Body)

	return err
}
