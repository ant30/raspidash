package web

import (
  "log"
  "html/template"
  "net/http"
  "strings"
)

var (
  templateMap = template.FuncMap{
    "Upper": func(s string) string {
      return strings.ToUpper(s)
    },
  }
  templates = template.New("").Funcs(templateMap)
)

// Parse all of the bindata templates
func init() {
  var s string
  for _, path := range AssetNames() {
    s = s + path
  }
  log.Println(s)
  for _, path := range AssetNames() {
    bytes, err := Asset(path)
    if err != nil {
      log.Panicf("Unable to parse: path=%s, err=%s", path, err)
    }
    templates.New(path).Parse(string(bytes))
  }
}

// The exposed method to render templates
func RenderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
  err := templates.ExecuteTemplate(w, tmpl, p)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
