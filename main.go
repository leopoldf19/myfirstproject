package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templ *template.Template

func init() {
	templ = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "index.gohtml", nil)
}

func registriertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	vorname := r.FormValue("vn")
	nachname := r.FormValue("nn")

	data := struct {
		Vorn  string
		Nachn string
	}{
		Vorn:  vorname,
		Nachn: nachname,
	}

	fmt.Println(data)

	templ.ExecuteTemplate(w, "registriert.gohtml", data)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/registriert", registriertHandler)
	http.ListenAndServe(":8080", nil)
}
