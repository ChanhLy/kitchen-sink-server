package router

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	database "go-server/db"
	"go-server/users"
	"go-server/utils"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func GetHandlers() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Heartbeat("/ping"))
	router.Use(middleware.Timeout(60 * time.Second))

	router.Mount("/", users.UserHandlers())

	return router
}

func ListenAndServe(handler http.Handler) {
	server := &http.Server{
		Addr:           utils.GetConfig().Port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	idleConnectionsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		log.Println("We received an interrupt signal, shut down.")
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		database.CloseDb()
		close(idleConnectionsClosed)
	}()

	l, _ := net.Listen("tcp", utils.GetConfig().Port)

	log.Println("Listening on " + utils.GetConfig().Port)

	if err := server.Serve(l); !errors.Is(err, http.ErrServerClosed) {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnectionsClosed
}
