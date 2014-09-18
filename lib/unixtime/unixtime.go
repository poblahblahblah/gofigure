// purpose: get the current time in unix time

// supports: CentOS, Darwin, Debian, Fedora, RedHat, Scientific, Ubuntu

// issues: 

package unixtime

import (
  "time"
  "strconv"
)

func Load() string {
  time_int64       := time.Now().Unix()
  time_string      := strconv.FormatInt(time_int64, 10)
  //time_string, err := strconv.ParseInt(time_int64, 10, 64)
  return time_string
}

