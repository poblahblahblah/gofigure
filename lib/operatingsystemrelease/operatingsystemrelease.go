package operatingsystemrelease

import (
  "github.com/poblahblahblah/gofigure/lib/factfuncts"
  "github.com/poblahblahblah/gofigure/lib/operatingsystem"
  "io/ioutil"
  "os"
  "strings"
)

func Load() string {
  switch operatingsystem.Load() {

  case "CentOS":
    return "FIXME"

  case "Debian":
    // debian is pretty straight forward. we should just be able to read the
    // contents of /etc/debian_version and be done with it.
    if _, err := os.Stat("/etc/debian_version"); err == nil {

      results, err := ioutil.ReadFile("/etc/debian_version")

      if err != nil {
        panic(err)
      }
      
      os_version := factfuncts.Chomp(string(results))

      return string(os_version)
    }

  case "Fedora":
    return "FIXME"

  case "Ubuntu":
    if _, err := os.Stat("/etc/lsb-release"); err == nil {

      results, err := ioutil.ReadFile("/etc/lsb-release")
      if err != nil { panic(err) }

      // scan through results w/ strings.Split
      for _, line :=  range strings.Split(string(results), "\n") {
        s := strings.Split(line, "=")
        if s[0] == "DISTRIB_RELEASE" {
          return factfuncts.Chomp(s[1])
        }
      }
    }
  }
  return "unsupported os"
}

