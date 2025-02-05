package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-server/users"
	"log"
	"net"
	"net/http"
	"time"
)

func Listen() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Timeout(60 * time.Second))

	r.Mount("/", users.UserHandlers())

	listenAndServe(r)

	return r
}

func listenAndServe(r *chi.Mux) {
	l, errListen := net.Listen("tcp", ":8080")
	if errListen != nil {
		log.Fatal(errListen)
	}

	log.Println("Listening on", l.Addr())

	err := http.Serve(l, r)
	if err != nil {
		log.Fatal(err)
	}
}
