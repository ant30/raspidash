package web

import (
	"log"
    "html/template"
	"net/http"
)

func renderIndex(w http.ResponseWriter) {
	t, err := template.ParseFiles("templates/layout.html", "templates/index.html")
    if err != nil {
        log.Println("Error loading the templates")
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    err2 := t.Execute(w, nil)
    if err2 != nil {
        log.Println("Error rendering the templates")
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderIndex(w)
}

func WebServer() {
	log.Println("Listening...")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
