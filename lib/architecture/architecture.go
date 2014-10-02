package architecture

import (
	"github.com/poblahblahblah/gofigure/lib/hardwaremodel"
	"github.com/poblahblahblah/gofigure/lib/operatingsystem"
	"regexp"
)

func Load() string {

	// regex so we can return i386 for all i686, etc.
	// this switch case statement is slightly different due to the regexes.
	// I couldn't figure out how to do it any other way, but it works.
	i386_regexp := regexp.MustCompile(`i[3456]86|pentium`)
	x86_64_regexp := regexp.MustCompile(`^x86_64$`)
	hardware := string(hardwaremodel.Load())

	switch {
	case x86_64_regexp.MatchString(hardware):
		switch operatingsystem.Load() {
		case "Debian", "Gentoo", "Ubuntu":
			return string("amd64")
		default:
			return string("x86_64")
		}

	case i386_regexp.MatchString(hardware):
		return string("i386")
	}

	return "unknown"

}
