package main

import (
    "github.com/kelseyhightower/go-facter"
    "github.com/poblahblahblah/gofigure/lib/fqdn"
    "github.com/poblahblahblah/gofigure/lib/hardwarearchitecture"
    "github.com/poblahblahblah/gofigure/lib/hostname"
    "github.com/poblahblahblah/gofigure/lib/kernelname"
    "github.com/poblahblahblah/gofigure/lib/kernelversion"
    "github.com/poblahblahblah/gofigure/lib/osarchitecture"
    "github.com/poblahblahblah/gofigure/lib/operatingsystem"
    "github.com/poblahblahblah/gofigure/lib/unixtime"
)

func main() {

    f := facter.New()
    f.Add("gofigure_version", gofigureVersion())
    f.Add("fqdn", fqdn.Load())
    f.Add("hardware_architecture", hardwarearchitecture.Load())
    f.Add("hostname", hostname.Load())
    f.Add("kernel", kernelname.Load())
    f.Add("kernel_version", kernelversion.Load())
    f.Add("operating_system", operatingsystem.Load())
    f.Add("os_architecture", osarchitecture.Load())
    f.Add("current_unix_time", unixtime.Load())
    
    f.Print()
}

func gofigureVersion() string {
  return string("0.0.1")
}

