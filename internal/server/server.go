package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nivasah/signet/internal/bookmark"
	"github.com/nivasah/signet/internal/pgdb"
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
	repo, err := pgdb.NewBookmarkRepository()
	if err != nil {
		log.Fatalf("no repository created: %v", err)
	}
	svc := bookmark.NewService(repo)
	setupRoutes(r, svc)
	r.Handle("/*", ui.SPAHandler())

	addr := fmt.Sprintf(":%s", cfg.Port)
	fmt.Println("Listening on", addr)
	http.ListenAndServe(addr, r)
	return nil
}

func setupRoutes(m *chi.Mux, svc bookmark.Service) {
	m.Route("/signet/api/v1", func(v1Route chi.Router) {
		v1Route.Route("/bookmarks", func(r chi.Router) {
			r.Post("/", func(w http.ResponseWriter, r *http.Request) {
				b := bookmark.Bookmark{
					Name:        "Simple Calculator Tutorial",
					Location:    "http://blog.dalw.in/simple-calculator-tutorial",
					Description: "This is a simple calculator tutorial, which I can use to learn programming",
					Tags:        "programming,simple,go,golang",
					FolderID:    0,
					UserID:      0,
				}

				err := svc.CreateBookmark(&b)
				if err != nil {
					log.Printf("unable to creat bookmark: %v", err)
					http.Error(w, "unable to create bookmark", http.StatusInternalServerError)
					return
				}
				w.Write([]byte("Successfully created bookmark"))
			})
		})
	})
}
