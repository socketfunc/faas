package internal

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func healthzHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		}
	}
}

func ListenAndServe(addr string) {
	http.HandleFunc("/_sf/healthz", healthzHandler())

	s := &http.Server{
		Addr: addr,
	}

	closed := make(chan struct{})
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGTERM)

		<-sig

		log.Println("SIGTERM received... shutting down server")

		if err := s.Shutdown(context.Background()); err != nil {
			log.Printf("Error in Shutdown: %v", err)
		}

		<-time.Tick(time.Duration(3) * time.Second)

		close(closed)
	}()
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("Error ListenAndServe: %v\n", err)
		close(closed)
	}

	<-closed
}
