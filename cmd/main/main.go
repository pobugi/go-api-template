package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/pobugi/go-crud-mysql/docs"
	"github.com/pobugi/go-crud-mysql/pkg/app"
	"github.com/pobugi/go-crud-mysql/pkg/books"
	"github.com/pobugi/go-crud-mysql/pkg/config"
	"github.com/pobugi/go-crud-mysql/pkg/db"
)

// @title           Bookstore API
// @version         1.0
// @description     CRUD demo with Gorilla Mux, GORM, MySQL.
// @BasePath        /
// @schemes         http

func main() {
	cfg := config.Load()
	gdb := db.Open(cfg)
	repo := books.NewRepo(gdb)

	srv := &http.Server{
		Addr:         cfg.HTTPAddr,
		Handler:      app.NewRouter(repo),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	go func() {
		log.Printf("listening on %s", cfg.HTTPAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	// graceful shutdown
	stop, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()
	<-stop.Done()

	ctx, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()
	_ = srv.Shutdown(ctx)
	log.Println("server stopped")
}
