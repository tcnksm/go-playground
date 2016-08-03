package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", withAuth(testHandler))

	if err := http.ListenAndServe(":3333", nil); err != nil {
		log.Fatal(err)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	if token := r.Context().Value(tokenKey); token != nil {
		w.Write([]byte("with Auth"))
		return
	}

	w.Write([]byte("without Auth"))
	return
}

const tokenKey = "tokenKey"

func withAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := "xxx"
		token := authorize(auth)
		ctx := context.WithValue(r.Context(), tokenKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func authorize(auth string) string {
	return "xxxxxxxxx"
}
