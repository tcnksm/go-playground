package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gorilla/websocket"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(filepath.Join("tmpl", "index.html.tmpl"))
		if err != nil {
			log.Printf("[ERROR] %s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err := t.Execute(w, nil); err != nil {
			log.Printf("[ERROR] %s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/socket", handleWebSocket)
	port := "3000"
	log.Printf("[INFO] Listening: %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Printf("[ERROR] %s", err)
	}
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(*http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")

	ws, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[ERROR] %s", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	log.Printf("[INFO] Connection open: %s", ws.RemoteAddr())

	message := []string{"A", "B", "C"}
	for _, m := range message {
		if err := ws.WriteMessage(websocket.TextMessage, []byte(m)); err != nil {
			log.Printf("[ERROR] %s", err)
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		time.Sleep(5 * time.Second)
	}
}
