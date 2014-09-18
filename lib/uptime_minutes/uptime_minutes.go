// purpose: returns a nodes uptime in minutes

// supports: Fedora

// issues: 

package uptime_minutes

import (
  "github.com/poblahblahblah/gofigure/lib/uptime_seconds"
  "strconv"
)

func Load() string {
  uptime_in_seconds := uptime_seconds.Load()
  uptimeInt, err    := strconv.Atoi(uptime_in_seconds)
  uptime_in_minutes := (uptimeInt / 60)
  if err != nil { return string("") }

  return strconv.Itoa(uptime_in_minutes)

}

