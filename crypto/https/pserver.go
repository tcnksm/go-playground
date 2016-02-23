package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", rootHandler)

	port := "3000"
	log.Printf("[INFO] Start listen on %s", port)
	err := http.ListenAndServe("localhost:"+port, nil)

	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO] Request from %s", r.RemoteAddr)
	w.Header().Set("Content-type", "text/plain")
	w.Write([]byte("Hello with TLS\n"))
}
