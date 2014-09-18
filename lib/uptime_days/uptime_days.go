// purpose: returns a nodes uptime in days

// supports: Fedora

// issues: 

package uptime_days

import (
  "github.com/poblahblahblah/gofigure/lib/uptime_seconds"
  "strconv"
)

func Load() string {
  uptime_in_seconds := uptime_seconds.Load()
  uptimeInt, err    := strconv.Atoi(uptime_in_seconds)
  uptime_in_days   := (uptimeInt / 60 / 60 / 24)
  if err != nil { return string("") }

  return strconv.Itoa(uptime_in_days)

}

