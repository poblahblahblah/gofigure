package hardwarearchitecture

import (
  "github.com/poblahblahblah/gofigure/lib/kernelname"
)

func Load() string {
// FIXME - no idea how to get this info from OS X.
    switch kernelname.Load() {
    case "Linux":
        return "x86_64"

    case "Darwin":
        return "i386"
    }
    return "unknown"
}

