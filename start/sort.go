/*
 * ----------------------------------------------------------------------------
 * "THE BEER-WARE LICENSE" (Revision 42):
 * <yazgazan@gmail.com> wrote this file. As long as you retain this notice you
 * can do whatever you want with this stuff. If we meet some day, and you think
 * this stuff is worth it, you can buy me a beer in return.
 * ----------------------------------------------------------------------------
 */

package start

import (
	"strings"
)

type dateSortFiles []indexItem

func (a dateSortFiles) Len() int {
	return len(a)
}

func (a dateSortFiles) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a dateSortFiles) Less(i int, j int) bool {
	if a[i].IsDir != a[j].IsDir { // folders first
		return a[i].IsDir
	}
	// most recent first
	return a[i].ModTime > a[j].ModTime
}

type alphaSortFiles []indexItem

func (a alphaSortFiles) Len() int {
	return len(a)
}

func (a alphaSortFiles) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a alphaSortFiles) Less(i int, j int) bool {
	if a[i].IsDir != a[j].IsDir { // folders first
		return a[i].IsDir
	}
	return strings.ToLower(a[i].Name) < strings.ToLower(a[j].Name)
}

type sizeSortFiles []indexItem

func (a sizeSortFiles) Len() int {
	return len(a)
}

func (a sizeSortFiles) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a sizeSortFiles) Less(i int, j int) bool {
	if a[i].IsDir != a[j].IsDir { // folders first
		return a[i].IsDir
	}
	// biggest file first
	return a[i].Size > a[j].Size
}
