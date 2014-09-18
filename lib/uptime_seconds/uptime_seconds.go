// purpose: returns a nodes uptime in seconds

// supports: Darwin, Fedora 

// issues:
//   * add "other" OSes

package uptime_seconds

import (
  "github.com/poblahblahblah/gofigure/lib/kernel"
  "github.com/poblahblahblah/gofigure/lib/unixtime"
  "os"
  "strings"
  "io/ioutil"
  "os/exec"
  "regexp"
  "strconv"
)

func Load() string {

  switch kernel.Load() {
  case "Linux":
    if _, err := os.Stat("/proc/uptime"); err == nil {
      proc_uptime, proc_uptime_err := ioutil.ReadFile("/proc/uptime")
      if proc_uptime_err != nil { panic("error") }
      return string(strings.Split(string(strings.Split(string(proc_uptime), " ")[0]), ".")[0])
    }

    return "linux failed"

  case "Darwin":
    if _, err := os.Stat("/usr/sbin/sysctl"); err == nil {
      sysctl_app              := "/usr/sbin/sysctl"
      sysctl_arg0             := "-a"
      sysctl_arg1             := "kern.boottime"
      sysctl_cmd              := exec.Command(sysctl_app, sysctl_arg0, sysctl_arg1)
      sysctl_out, sysctl_err  := sysctl_cmd.Output()

      if sysctl_err != nil { panic("error") }

      sysctl_regexp := regexp.MustCompile(`sec = (\d+)`)

      current_time_int64, ct_err := strconv.ParseInt(strings.Split(sysctl_regexp.FindString(string(sysctl_out)), "sec = ")[1], 10, 64)
      boot_time_int64, bt_err    := strconv.ParseInt(unixtime.Load(), 10, 64)
      uptime_seconds_int64       := (boot_time_int64 - current_time_int64)

      if ct_err != nil { panic(ct_err) }
      if bt_err != nil { panic(bt_err) }

      return strconv.FormatInt(uptime_seconds_int64, 10)

    }

    return "darwin failed"
    
  }
  return string("FIXME")
}

