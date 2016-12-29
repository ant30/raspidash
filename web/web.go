package web

import (
  "flag"
  "log"
  "net/http"

)

var (
    serverCommand = flag.NewFlagSet("server", flag.ExitOnError)
)

func WebServer(args []string) {
    serverListen := serverCommand.String("listen", ":8000", "This is the tcp port to bind the http server interface, like :8000 or 0.0.0.0:8000")
    serverCommand.Parse(args)
    log.Printf("Preparing to listen on %v \n", *serverListen)


    http.HandleFunc("/", IndexHandler)
    http.HandleFunc("/display/", DisplayHandler)
    log.Println("Listening...")
    log.Println(http.ListenAndServe(*serverListen, nil))
}
