package web

import (
  "net/http"
  "log"
)

type displayContext struct {
  Name string
  Url string
}

func DisplayHandler(w http.ResponseWriter, r *http.Request) {
  log.Printf("%s: %s\n", r.Method, r.URL.Path)
  if r.Method == "GET" {
	RenderTemplate(w, "templates/display-view.html", displayContext{"MainScreen", "http://www.anyweb.com/"})
  } else {
	return
  }
}

