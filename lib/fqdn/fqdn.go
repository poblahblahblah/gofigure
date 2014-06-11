package fqdn

import (
  "github.com/poblahblahblah/gofigure/lib/factfuncts"
  "os/exec"
)

func Load() string {
  app  := "hostname"
  arg0 := "-f"

  cmd := exec.Command(app, arg0)
  out, err := cmd.Output()

  if err != nil {
    panic(err)
  }

  return factfuncts.Chomp(string(out))
}

