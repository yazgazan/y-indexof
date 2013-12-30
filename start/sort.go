
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

