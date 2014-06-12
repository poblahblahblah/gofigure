package path

import (
  "github.com/poblahblahblah/gofigure/lib/factfuncts"
  "os"
)

func Load() string {
  return factfuncts.Chomp(os.Getenv("PATH"))
}

