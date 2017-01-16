package web

import (
  "net/http"
  "log"
  "github.com/ant30/raspidash/discovery"
)

type indexContext struct {
  Title string
  DeviceList []discovery.ServiceDescription
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  log.Printf("%s: %s\n", r.Method, r.URL.Path)
  if r.Method == "GET" {
    RenderTemplate(w, "templates/index.html", indexContext{
        Title: "So what!",
        DeviceList: discovery.ListDevices(),
    })
  } else {
	return
  }
}
