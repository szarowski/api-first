package api

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	*http.Server
}

func NewStrictServer(r *chi.Mux) *Server {
	srv := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &Server{&srv}
}

func (srv *Server) Start() {
	log.Println("Starting Chi HTTP Server...")
	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
	log.Printf("Listening on %s\n", srv.Addr)

	srv.waitForShutdown()
}

func (srv *Server) waitForShutdown() {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal
	<-interruptChan

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server shutdown failed due to: %v", err)
	}

	log.Println("Server gracefully shut down")
	os.Exit(0)
}
