package web

import (
    "flag"
    "log"
    "net/http"
    "github.com/ant30/raspidash/discovery"
)

var (
    serverCommand = flag.NewFlagSet("server", flag.ExitOnError)
    discoverable discovery.Service
)

func WebServer(args []string) {
    LoadBinTemplates()

    serverListen := serverCommand.String("listen", ":8000", "This is the tcp port to bind the http server interface, like :8000 or 0.0.0.0:8000")
    serverName := serverCommand.String("name", discovery.GetHostname(), "This is the device name, using current hostname if not define not defined")
    serverCommand.Parse(args)
    log.Printf("Preparing device %v to listen on %v \n", *serverName, *serverListen)

    log.Println("Doing this service discoverable")
    go discoverable.DoDiscoverable(*serverName, *serverListen)

    http.HandleFunc("/", IndexHandler)
    http.HandleFunc("/display/", DisplayHandler)
    log.Println("Listening...")
    log.Println(http.ListenAndServe(*serverListen, nil))
}
