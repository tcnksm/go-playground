package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func main() {

	http.HandleFunc("/text", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("static", "text.html"))
	})

	http.HandleFunc("/textarea", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("static", "textarea.html"))
	})

	http.HandleFunc("/radio", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("static", "radio.html"))
	})

	http.HandleFunc("/checkbox", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("static", "checkbox.html"))
	})

	http.HandleFunc("/input/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%#v", r)
		segs := strings.Split(r.URL.Path, "/")
		action := segs[2]
		switch action {
		case "text":
			firstname := r.FormValue("firstname")
			lastname := r.FormValue("lastname")
			fmt.Fprintf(w, "firstname is '%s', lastname is '%s'\n", firstname, lastname)

		case "textarea":
			memo := r.FormValue("memo")
			fmt.Fprintf(w, "memo is '%s'\n", memo)

		case "radio":
			gender := r.FormValue("gender")
			fmt.Fprintf(w, "gendor is '%s'\n", gender)

		case "checkbox":
			r.ParseForm()
			interest := r.Form["interest"]
			fmt.Fprintf(w, "gendor is '%#v'\n", interest)

		default:
			log.Printf("Invalid action: %s", action)
		}
	})

	log.Printf("Start listening :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
