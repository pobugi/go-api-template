package health

import (
	"net/http"

	"github.com/pobugi/go-crud-mysql/pkg/httpx"
)

func Live() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}
}

func Ready() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.JSON(w, http.StatusOK, map[string]string{"status": "ready"})
	}
}
