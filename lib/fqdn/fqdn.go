package fqdn

import (
  "github.com/poblahblahblah/gofigure/lib/domain"
  "github.com/poblahblahblah/gofigure/lib/hostname"
  "strings"
)

func Load() string {
  // this will take the data from the `hostname` and `domain` facts and throw
  // them together. this is to (hopefully) avoid a DNS lookup if possible.
  // this could all be one line, but separated it out for clarity.
  fqdn_bits := []string{hostname.Load(), domain.Load()}
  fqdn      := strings.Join(fqdn_bits, ".")

  return string(fqdn)

}

