package timezone

import (
  "github.com/poblahblahblah/gofigure/lib/factfuncts"
  "os/exec"
)

func Load() string {
  app  := "date"
  arg0 := "+%Z"

  cmd := exec.Command(app, arg0)
  // Ignore the current environment so that we are not influenced by a TZ variable
  // and instead get the system timezone
  cmd.Env = []string{}
  out, err := cmd.Output()

  if err != nil {
    return string("")
  }

  return factfuncts.Chomp(string(out))
}

