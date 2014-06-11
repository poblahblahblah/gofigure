package kernelmajversion

import (
  "github.com/poblahblahblah/gofigure/lib/kernelversion"
  "strings"
)

func Load() string {
  kernel := strings.Split(kernelversion.Load(), ".")
  k      := []string{kernel[0], kernel[1]}
  kernelmajversion := strings.Join(k, ".")

  return kernelmajversion
}

