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

type DateSortFiles []IndexItem

func (a DateSortFiles) Len() int {
  return len(a)
}

func (a DateSortFiles) Swap(i int, j int) {
  a[i], a[j] = a[j], a[i]
}

func (a DateSortFiles) Less(i int, j int) bool {
  if a[i].IsDir != a[j].IsDir == true { // folders first
    return a[i].IsDir
  }
  // most recent first
  return a[i].ModTime > a[j].ModTime
}

type AlphaSortFiles []IndexItem

func (a AlphaSortFiles) Len() int {
  return len(a)
}

func (a AlphaSortFiles) Swap(i int, j int) {
  a[i], a[j] = a[j], a[i]
}

func (a AlphaSortFiles) Less(i int, j int) bool {
  if a[i].IsDir != a[j].IsDir == true { // folders first
    return a[i].IsDir
  }
  return strings.ToLower(a[i].Name) < strings.ToLower(a[j].Name)
}

type SizeSortFiles []IndexItem

func (a SizeSortFiles) Len() int {
  return len(a)
}

func (a SizeSortFiles) Swap(i int, j int) {
  a[i], a[j] = a[j], a[i]
}

func (a SizeSortFiles) Less(i int, j int) bool {
  if a[i].IsDir != a[j].IsDir == true { // folders first
    return a[i].IsDir
  }
  // biggest file first
  return a[i].Size > a[j].Size
}

