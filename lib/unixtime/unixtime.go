package unixtime

import (
  "time"
)

func Load() string {
  return string(time.Now().Unix())
}

