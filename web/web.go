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
    log.Println("Doing this service discoverable")
    go discoverable.DoDiscoverable("tv1", 8000)

    serverListen := serverCommand.String("listen", ":8000", "This is the tcp port to bind the http server interface, like :8000 or 0.0.0.0:8000")
    serverCommand.Parse(args)
    log.Printf("Preparing to listen on %v \n", *serverListen)

    http.HandleFunc("/", IndexHandler)
    http.HandleFunc("/display/", DisplayHandler)
    log.Println("Listening...")
    log.Println(http.ListenAndServe(*serverListen, nil))
}
