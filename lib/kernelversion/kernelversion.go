package kernelversion

import (
  "github.com/poblahblahblah/gofigure/lib/kernelrelease"
  "strings"
)

func Load() string {
  kernelrelease := strings.Split(kernelrelease.Load(), "-")

  return kernelrelease[0]
}

