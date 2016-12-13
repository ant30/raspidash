package web

import (
  "net/http"
  "log"
)

type indexContext struct {
  Title string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  log.Printf("%s: %s\n", r.Method, r.URL.Path)
  if r.Method == "GET" {
    RenderTemplate(w, "templates/index.html", indexContext{"So what!"})
  } else {
	return
  }
}
