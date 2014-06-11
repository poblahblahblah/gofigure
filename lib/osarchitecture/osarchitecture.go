package osarchitecture

import (
  "github.com/poblahblahblah/gofigure/lib/factfuncts"
  "os/exec"
)

func Load() string {
    app  := "uname"
    arg0 := "-p"

    cmd := exec.Command(app, arg0)
    out, err := cmd.Output()

    if err != nil {
        panic(err)
    }
    return(factfuncts.Chomp(string(out)))
}

