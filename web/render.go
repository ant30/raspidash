package web

import (
  "log"
  "html/template"
  "net/http"
  "strings"
)

var (
  templateFuncMap = template.FuncMap{
    "Upper": func(s string) string {
      return strings.ToUpper(s)
    },
  }
  baseTemplate = template.New("").Funcs(templateFuncMap)
  templates map[string]*template.Template
)

// Parse all of the bindata templates
func init() {
  log.Printf("loading base layout")
  baseBytes, err := layoutsBaseHtmlBytes()
    if err != nil {
      log.Panicf("Unable to read base template: err=%s", err)
    }

  baseTemplate.New("layouts/base.html").Parse(string(baseBytes))
  templates = make(map[string]*template.Template)

  for _, path := range AssetNames() {
	if strings.HasPrefix("/layouts/", path) {
		continue
	}
	log.Printf("Loading template %s\n", path)
    bytes, err := Asset(path)
    if err != nil {
      log.Panicf("Unable to read template: path=%s, err=%s", path, err)
    }
	templates[path], err = baseTemplate.Clone()
	templates[path] = template.Must(templates[path].Parse(string(bytes)))
  }
}

// The exposed method to render templates
func RenderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {

  err := templates[tmpl].Execute(w, p)
  if err != nil {
	log.Panicf("Error during render: tmpl %s , error %s", tmpl, err.Error())
    http.Error(w, err.Error(), http.StatusInternalServerError)
	return
  }
}
