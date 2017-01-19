package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lestrrat/go-server-starter/listener"
	"github.com/shogo82148/go-gracedown"
)

func main() {

	listeners, err := listener.ListenAll()
	if err != nil || len(listeners) == 0 {
		log.Fatal("[ERROR] Failed to get a listener: ", err)
	}

	// subscribe to SIGINT signals
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGTERM)

	// go18(listeners[0], sigCh)
	goGracedown(listeners[0], sigCh)
	log.Println("[INFO] Gracefully shutdown")
}

func go18(l net.Listener, sigCh chan os.Signal) {
	log.Println("[INFO] Use go1.8")
	server := &http.Server{Handler: newHandler()}
	go func() {
		if err := server.Serve(l); err != nil {
			log.Println("[ERROR] Serve:", err)
		}
	}()

	<-sigCh
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("[INFO] Shutting down server")
	if err := server.Shutdown(ctx); err != nil {
		log.Println("[ERROR]", err)
	}

}

func goGracedown(l net.Listener, sigCh chan os.Signal) {
	log.Println("[INFO] Use go-gracedown")
	go func() {
		if err := gracedown.Serve(l, newHandler()); err != nil {
			log.Println("[ERROR]", err)
		}
	}()

	<-sigCh
	log.Println("[INFO] Shutting down server")
	gracedown.Close()
}

func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INFO] %s %s", r.Method, r.URL.Path)
		time.Sleep(100 * time.Millisecond)
	})
	return mux
}
