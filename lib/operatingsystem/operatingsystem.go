package operatingsystem

import (
  "github.com/poblahblahblah/gofigure/lib/factfuncts"
  "io/ioutil"
  "os"
  "regexp"
  "strings"
)

func Load() string {

  // try to detect which OS we are
  // if /etc/redhat-release exists then we're dealing with redhat, centos, or scientific
  if _, err := os.Stat("/etc/redhat-release"); err == nil {

    centos_regexp, err     := regexp.Compile(`(?i)centos`)
    fedora_regexp, err     := regexp.Compile(`(?i)fedora`)
    redhat_regexp, err     := regexp.Compile(`(?i)redhat`)
    scientific_regexp, err := regexp.Compile(`(?i)scientific`)

    results, err := ioutil.ReadFile("/etc/redhat-release")
    if err != nil {
      panic(err)
    }

    os_long_string := factfuncts.Chomp(string(results))

    if centos_regexp.MatchString(os_long_string) == true {
      return string("CentOS")
    } else if redhat_regexp.MatchString(os_long_string) == true {
      return string("RedHat")
    } else if scientific_regexp.MatchString(os_long_string) == true {
      return string("Scientific Linux")
    }
  }

	
  // if /etc/debian_release exists, then we're using debian or a derivative
  if _, err := os.Stat("/etc/debian_release"); err == nil {

    // if /etc/lsb-release
    if _, err := os.Stat("/etc/lsb-release"); err == nil {

      // Read in /etc/lsb-release
      results, err := ioutil.ReadFile("/etc/lsb-release")
      if err != nil {
        panic(err)
      }

      // scan through results w/ strings.Split
      for _, line :=  range strings.Split(string(results), "\n") {
        s := strings.Split(line, "=")
        if s[0] == "DISTRIB_DESCRIPTION" {
          return(strings.Replace(s[1], `"`, "", -1))
        }
      }
    }
  }

  // FIXME:
  // Add Solaris, the BSDs, etc.

  return(string("error: cannot detect operating system"))
}

