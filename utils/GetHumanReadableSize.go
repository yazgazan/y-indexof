
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

