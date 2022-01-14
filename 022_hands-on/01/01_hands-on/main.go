package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}
func main() {
	http.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(rw, "template.gohtml", "Hello world")
	}))
	http.HandleFunc("/dog", func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "Dogo")
	})
	http.HandleFunc("/me", func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "Apaar")
	})
	http.ListenAndServe(":8080", nil)
}
