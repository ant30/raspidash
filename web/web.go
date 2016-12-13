package web

import (
  "log"
  "net/http"
)


func WebServer() {
  http.HandleFunc("/", IndexHandler)
  http.HandleFunc("/display/", DisplayHandler)
  log.Println("Listening...")
  log.Println(http.ListenAndServe(":8000", nil))
}
