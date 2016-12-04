package web

import (
	"io"
	"log"
	"net/http"
)

func renderIndex(w http.ResponseWriter) {
	io.WriteString(w, "Simple Index page for raspidash project")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderIndex(w)
}

func WebServer() {
	log.Println("Listening...")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
