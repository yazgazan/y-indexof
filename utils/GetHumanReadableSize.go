/*
 * ----------------------------------------------------------------------------
 * "THE BEER-WARE LICENSE" (Revision 42):
 * <yazgazan@gmail.com> wrote this file. As long as you retain this notice you
 * can do whatever you want with this stuff. If we meet some day, and you think
 * this stuff is worth it, you can buy me a beer in return.
 * ----------------------------------------------------------------------------
 */

package utils

import (
  "fmt"
)

func GetHumanReadableSize(size int64) string {
  switch {
    case size < 1024: return fmt.Sprintf("%dB", size)
    case size < 1024 * 1024: return fmt.Sprintf("%dkB", size / 1024)
    case size < 1024 * 1024 * 1024: return fmt.Sprintf("%dMB", size / (1024 * 1024))
    case size < 1024 * 1024 * 1024 * 1024:
      return fmt.Sprintf("%dGB", size / (1024 * 1024 * 1024))
    default: return fmt.Sprintf("%dB", size)
  }
  return ""
}

