// purpose: get a node's domain

// supports: CentOS, Darwin, Debian, Fedora, RedHat, Scientific, Ubuntu

// issues: returning an empty string is not ideal.

package domain

import (
  "github.com/poblahblahblah/gofigure/lib/factfuncts"
  "os/exec"
  "regexp"
)

func Load() string {

  // first we'll do a check to see if we can just use the hostname provided by
  // `hostname` to pull the domain out. This would help us avoid using the `-f`
  // argument for `hostname` which causes DNS lookups on some platforms.

  hostname_app               := "hostname"
  hostname_arg0              := "-s"
  hostname_cmd               := exec.Command(hostname_app, hostname_arg0)
  hostname_out, hostname_err := hostname_cmd.Output()
  hostname                   := factfuncts.Chomp(string(hostname_out))

  if hostname_err !=nil { panic(hostname_err) }

  domain_regexp := regexp.MustCompile(`.*?\.(.+$)`)

  if domain_regexp.MatchString(hostname) {
    domain := domain_regexp.FindStringSubmatch(hostname)
    return string(domain[1])
  }

  // if we're here that means that hostname returned a short name.
  // run `hostname -f` so we can pick out the domain.

  hostname_f_arg0                := "-f"
  hostname_f_cmd                 := exec.Command(hostname_app, hostname_f_arg0)
  hostname_f_out, hostname_f_err := hostname_f_cmd.Output()
  hostname_f                     := factfuncts.Chomp(string(hostname_f_out))

  if hostname_f_err != nil { 
    // if we get here that means hostname -f failed.
    // we should return nothing if that is the case.
    return string("")
  }

  if domain_regexp.MatchString(hostname_f) {
    domain := domain_regexp.FindStringSubmatch(hostname_f)
    return string(domain[1])
  }

  //  I don't know why we would be here.
  return "FIXME"
}

