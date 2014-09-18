package operatingsystemrelease

import (
  "github.com/poblahblahblah/gofigure/lib/factfuncts"
  "github.com/poblahblahblah/gofigure/lib/operatingsystem"
  "io/ioutil"
  "os"
  "regexp"
  "strings"
)

func Load() string {
  switch operatingsystem.Load() {

  case "CentOS", "RedHat", "Scientific":
    // FIXME: there's probably a better way to do this.
    if _, err := os.Stat("/etc/redhat-release"); err == nil {

      results, err := ioutil.ReadFile("/etc/redhat-release")
      if err !=nil { panic(err) }

      version_regexp, err := regexp.Compile(`\d+\.\d+`)
      return version_regexp.FindString(string(results))
    }


  case "Debian":
    // debian is pretty straight forward. we should just be able to read the
    // contents of /etc/debian_version and be done with it.
    if _, err := os.Stat("/etc/debian_version"); err == nil {

      results, err := ioutil.ReadFile("/etc/debian_version")
      if err != nil { panic(err) }
      
      os_version := factfuncts.Chomp(string(results))
      return string(os_version)
    }

  case "Fedora":
    // FIXME: this is also probably terrible.
    if _, err := os.Stat("/etc/redhat-release"); err == nil {

      results, err := ioutil.ReadFile("/etc/redhat-release")
      if err !=nil { panic(err) }

      version_regexp, err := regexp.Compile(`\d+`)
      return (version_regexp.FindString(string(results)))
    }

  case "OracleLinux":
    // FIXME: I feel like this will break pretty easily.
    if _, err := os.Stat("/etc/oracle-release"); err == nil {

    results, err := ioutil.ReadFile("/etc/oracle-release")
    if err != nil { panic(err) }

    version_regexp, err := regexp.Compile(`\d+\.\d+`)
    return (version_regexp.FindString(string(results)))

  }


  case "Darwin":
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

