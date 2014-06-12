package main

import (
    "github.com/kelseyhightower/go-facter"
    "github.com/poblahblahblah/gofigure/lib/architecture"
    "github.com/poblahblahblah/gofigure/lib/fqdn"
    "github.com/poblahblahblah/gofigure/lib/hardwareisa"
    "github.com/poblahblahblah/gofigure/lib/hardwaremodel"
    "github.com/poblahblahblah/gofigure/lib/hostname"
    "github.com/poblahblahblah/gofigure/lib/id"
    "github.com/poblahblahblah/gofigure/lib/kernel"
    "github.com/poblahblahblah/gofigure/lib/kernelmajversion"
    "github.com/poblahblahblah/gofigure/lib/kernelrelease"
    "github.com/poblahblahblah/gofigure/lib/kernelversion"
    "github.com/poblahblahblah/gofigure/lib/operatingsystem"
    "github.com/poblahblahblah/gofigure/lib/operatingsystemrelease"
    "github.com/poblahblahblah/gofigure/lib/osarchitecture"
    "github.com/poblahblahblah/gofigure/lib/osfamily"
    "github.com/poblahblahblah/gofigure/lib/path"
    "github.com/poblahblahblah/gofigure/lib/timezone"
    "github.com/poblahblahblah/gofigure/lib/unixtime"
)

func main() {

    f := facter.New()
    f.Add("gofigure_version",       gofigureVersion())
    f.Add("architecture",           architecture.Load())
    f.Add("fqdn",                   fqdn.Load())
    f.Add("hardwareisa",            hardwareisa.Load())
    f.Add("hardwaremodel",          hardwaremodel.Load())
    f.Add("hostname",               hostname.Load())
    f.Add("id",                     id.Load())
    f.Add("kernel",                 kernel.Load())
    f.Add("kernelversion",          kernelmajversion.Load())
    f.Add("kernelmajversion",       kernelmajversion.Load())
    f.Add("kernelrelease",          kernelrelease.Load())
    f.Add("kernelversion",          kernelversion.Load())
    f.Add("operatingsystem",        operatingsystem.Load())
    f.Add("operatingsystemversion", operatingsystemrelease.Load())
    f.Add("osfamily",               osfamily.Load())
    f.Add("os_architecture",        osarchitecture.Load())
    f.Add("path",                   path.Load())
    f.Add("timezone",               timezone.Load())
    f.Add("current_unix_time",      unixtime.Load())
    
    f.Print()
}

func gofigureVersion() string {
  return string("0.0.1")
}

