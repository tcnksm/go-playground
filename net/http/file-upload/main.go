package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		if method != "POST" {
			msg := fmt.Sprintf("Method '%s' is not allowed", method)
			w.Write([]byte(msg))
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		file, header, err := r.FormFile("uploading")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("Failed: %s", err)))
			return
		}
		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("Failed: %s", err)))
			return
		}

		filename := "uploaded" + filepath.Ext(header.Filename)
		if err := ioutil.WriteFile(filename, data, 0600); err != nil {
			w.Write([]byte(fmt.Sprintf("Failed: %s", err)))
			return
		}

		w.Write([]byte("Success\n"))
	})

	log.Printf("Start listening :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
