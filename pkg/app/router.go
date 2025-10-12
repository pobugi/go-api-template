package app

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/pobugi/go-crud-mysql/pkg/books"
	"github.com/pobugi/go-crud-mysql/pkg/health"
	"github.com/pobugi/go-crud-mysql/pkg/httpx"
	httpSwagger "github.com/swaggo/http-swagger"
)

// NewRouter wires middlewares and routes.
func NewRouter(repo *books.Repo) http.Handler {
	v := validator.New()
	r := mux.NewRouter()

	// global middleware
	r.Use(httpx.Recover, httpx.Logging, httpx.RequestID, httpx.JSONContentType)

	// docs
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler).Methods("GET")

	// health
	r.HandleFunc("/health", health.Live()).Methods("GET")
	r.HandleFunc("/ready", health.Ready()).Methods("GET")

	// v1 api
	api := r.PathPrefix("/v1").Subrouter()
	api.HandleFunc("/books", books.Create(repo, v)).Methods("POST")
	api.HandleFunc("/books", books.List(repo)).Methods("GET")
	api.HandleFunc("/books/{id}", books.GetByID(repo)).Methods("GET")
	api.HandleFunc("/books/{id}", books.Update(repo, v)).Methods("PUT")
	api.HandleFunc("/books/{id}", books.Delete(repo)).Methods("DELETE")

	return r
}
