package unixtime

import (
  "time"
)

func Load() int64 {
  return time.Now().Unix()
}

