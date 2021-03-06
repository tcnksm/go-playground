package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	// Generated by openssl command
	certFile, _ := filepath.Abs("certificate/server.crt")
	keyFile, _ := filepath.Abs("certificate/server.key")

	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))),
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// https://go-review.googlesource.com/#/c/32412/
		// https://blog.tylerchr.com/golang-18-whats-coming/

		if p, ok := w.(http.Pusher); ok {
			p.Push("/static/gopher.jpg", nil)
		}

		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte(`<img src="/static/gopher.jpg"/>`))
	})

	log.Println(":3000")
	if err := http.ListenAndServeTLS(":3000", certFile, keyFile, nil); err != nil {
		log.Printf("[ERROR] %s", err)
	}
}
