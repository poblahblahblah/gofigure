package osfamily

import (
  "github.com/poblahblahblah/gofigure/lib/operatingsystem"
)

func Load() string {
  // let's try to work through the operating systems alphabetically.

  switch operatingsystem.Load() {
  case "CentOS", "Fedora", "OracleLinux", "RedHat", "Scientific":
    return "RedHat"
  case "Debian", "Ubuntu":
    return "Debian"
  }
  return "unknown"

}

