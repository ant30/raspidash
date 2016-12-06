package web

import (
  "log"
  "net/http"
)


func WebServer() {
  log.Println("Listening...")
  http.HandleFunc("/", IndexHandler)
  http.ListenAndServe(":8000", nil)
}
