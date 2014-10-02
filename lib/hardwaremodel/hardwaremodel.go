// purpose: returns hardware model of the current node

// supports: CentOS, Darwin, Debian, Fedora, RedHat, Scientific, Ubuntu

package hardwaremodel

import (
	"github.com/poblahblahblah/gofigure/lib/factfuncts"
	"os/exec"
)

func Load() string {
	app := "uname"
	arg0 := "-m"
	cmd := exec.Command(app, arg0)
	out, err := cmd.Output()

	if err != nil {
		return string("")
	}

	return factfuncts.Chomp(string(out))
}
