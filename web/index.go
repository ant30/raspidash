package web

import (
  "net/http"
)

type indexContext struct {
  Title string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == 'GET' {
    RenderTemplate(w, "templates/index.html", indexContext{"So what!"})
  }
}
