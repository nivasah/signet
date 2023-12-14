package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nivasah/signet/ui"
)

type Config struct {
	Port string
}

func Run(cfg *Config) error {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})
	r.Handle("/*", ui.SPAHandler())

	addr := fmt.Sprintf(":%s", cfg.Port)
	fmt.Println("Listening on", addr)
	http.ListenAndServe(addr, r)
	return nil
}
