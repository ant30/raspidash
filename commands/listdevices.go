package commands

import (
    "flag"
    "github.com/ant30/raspidash/discovery"
    "fmt"
)

var (
    listDevicesCommand = flag.NewFlagSet("list", flag.ExitOnError)
)


func ListDevices(args []string) {
    services := discovery.ListDevices()
    for _, s := range services {
        fmt.Printf("%s\t%s\t%s\t%d\n", s.Name, s.Host, s.IPv4.String(), s.Port)
    }
}
