// purpose: pull out major release version for each operatingsystem.

// supports: CentOS, Debian, Fedora, RedHat, Scientific

package operatingsystemmajrelease

import (
  "github.com/poblahblahblah/gofigure/lib/operatingsystem"
  "github.com/poblahblahblah/gofigure/lib/operatingsystemrelease"
  "regexp"
)

func Load() string {
  switch operatingsystem.Load() {

  case "CentOS", "Debian", "RedHat", "Scientific":
    version_regexp, err := regexp.Compile(`\d+`)
    if err !=nil { panic(err) }

    return (version_regexp.FindString(operatingsystemrelease.Load()))

  case "Fedora", "Ubuntu":
    return operatingsystemrelease.Load()

  case "OracleLinux":
    return "FIXME"

  case "OS X":
    return "FIXME"

  }
  return "unsupported os"
}

