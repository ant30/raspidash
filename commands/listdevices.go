package commands

import (
    "log"
    "flag"
    "github.com/ant30/raspidash/discovery"
)

var (
    listDevicesCommand = flag.NewFlagSet("list", flag.ExitOnError)
)


func ListDevices(args []string) {
    services := discovery.ListDevices()
    log.Printf("Services:  %#v", services)
}
