// purpose: returns a nodes uptime (rounded down to the date and then to the hour)

// supports: Fedora

// issues:

package uptime_hours

import (
	"github.com/poblahblahblah/gofigure/lib/uptime_seconds"
	"strconv"
)

func Load() string {
	uptime_in_seconds := uptime_seconds.Load()
	uptimeInt, err := strconv.Atoi(uptime_in_seconds)
	uptime_in_hours := (uptimeInt / 60 / 60)
	if err != nil {
		return string("")
	}

	return strconv.Itoa(uptime_in_hours)

}
