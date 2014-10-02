// purpose: returns node's hardware processor type

// supports: CentOS, Darwin, Debian, Fedora, RedHat, Scientific, Ubuntu

// issues:

package hardwareisa

import (
	"github.com/poblahblahblah/gofigure/lib/factfuncts"
	"os/exec"
)

func Load() string {
	uname_app := "uname"
	uname_arg0 := "-p"
	uname_cmd := exec.Command(uname_app, uname_arg0)
	uname_out, uname_err := uname_cmd.Output()

	if uname_err != nil {
		return string("")
	}

	return factfuncts.Chomp(string(uname_out))
}
