package hardwareisa

import (
  "github.com/poblahblahblah/gofigure/lib/factfuncts"
  "os/exec"
)

func Load() string {
  app      := "uname"
  arg0     := "-p"
  cmd      := exec.Command(app, arg0)
  out, err := cmd.Output()

  if err != nil { return string("") }

  return factfuncts.Chomp(string(out))
}

