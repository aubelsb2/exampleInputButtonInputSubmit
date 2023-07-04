package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	//go:embed "index.html"
	indexHtmlData []byte
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/process", processHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("index.html").Parse(string(indexHtmlData)))
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		name := r.FormValue("name")
		action := r.FormValue("action")

		fmt.Fprintf(w, "Name: %s\n", name)
		fmt.Fprintf(w, "Action: %s\n", action)
	}
}
