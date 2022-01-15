package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}
func main() {
	http.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "foo ran")
	}))
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/assets/dog.jpg", dogPic)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/dog.jpg")
}
