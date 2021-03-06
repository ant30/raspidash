package main

import (
    "flag"
    "log"
    "os"
    "github.com/ant30/raspidash/web"
    "github.com/ant30/raspidash/commands"
)


func main() {
    if len(os.Args) < 2 {
        log.Println("a subcommand is required [server]")
        os.Exit(1)
    }
    switch os.Args[1] {
        case "server":
            web.WebServer(os.Args[2:])
        case "check":
            commands.CheckSyntax(os.Args[2:])
        case "list":
            commands.ListDevices(os.Args[2:])
        default:
            flag.PrintDefaults()
            os.Exit(1)
    }
}
