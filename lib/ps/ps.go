// purpose: returns default ps string (used by Puppet)

// supports: CentOS, Darwin, Debian, DragonFly, Fedora, FreeBSD, NetBSD,
//           OpenBSD, RedHat, Scientific, Ubuntu

// issues:

package ps

import (
  "github.com/poblahblahblah/gofigure/lib/operatingsystem"
)

func Load() string {

  switch operatingsystem.Load() {

  case "Darwin", "DragonFly", "FreeBSD", "NetBSD", "OpenBSD":
    return "ps auxwww"

  default:
    return "ps -ef"
  }

  return "unknown"

}

