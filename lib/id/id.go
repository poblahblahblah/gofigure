package id

import (
  "github.com/poblahblahblah/gofigure/lib/factfuncts"
  "github.com/poblahblahblah/gofigure/lib/kernel"
  "os/exec"
)

func Load() string {
  switch kernel.Load() {
  case "SunOS":
    app      := "/usr/xpg4/bin/id"
    arg0     := "-un"
    cmd      := exec.Command(app, arg0)
    out, err := cmd.Output()

    if err !=nil { return string("") }

    return factfuncts.Chomp(string(out))

  default:
    app      := "whoami"
    cmd      := exec.Command(app)
    out, err := cmd.Output()

    if err != nil { return string("") }

    return factfuncts.Chomp(string(out))
  }
  return "unsupported OS"
}

