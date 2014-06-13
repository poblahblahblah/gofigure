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

  app      := "hostname"
  cmd      := exec.Command(app)
  out, err := cmd.Output()
  hostname := factfuncts.Chomp(string(out))

  if err !=nil { panic(err) }

  domain_regexp := regexp.MustCompile(`.*?\.(.+$)`)

  switch {
    case domain_regexp.MatchString(hostname):
      domain := domain_regexp.FindStringSubmatch(hostname)
      return string(domain[1])

    default:
      app      := "hostname"
      arg0     := "-f"
      cmd      := exec.Command(app, arg0)
      out, err := cmd.Output()

      if err != nil { 
        // if we get here that means hostname -f failed.
        // we should return nothing if that is the case.
        return string("")
      }

      if domain_regexp.MatchString(string(out)) {
        domain := domain_regexp.FindStringSubmatch(string(out))
        return string(domain[1])
      }
  }

  // I guess this means that there's no domain set. weird.
  return ""
}

