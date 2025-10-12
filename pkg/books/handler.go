package books

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/pobugi/go-crud-mysql/pkg/httpx"
)

// @Summary      Create book
// @Description  Create a new book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        payload  body      CreateBookRequest  true  "book"
// @Success      201      {object}  httpx.Envelope{data=BookResponse}
// @Failure      400      {object}  httpx.Envelope
// @Failure      500      {object}  httpx.Envelope
// @Router       /v1/books [post]
func Create(repo *Repo, v *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var in CreateBookRequest
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			httpx.Fail(w, http.StatusBadRequest, "bad_json", "invalid json")
			return
		}
		if err := v.Struct(in); err != nil {
			httpx.Fail(w, http.StatusBadRequest, "validation", err.Error())
			return
		}
		b := &Book{Name: in.Name, Author: in.Author, Publication: in.Publication}
		if err := repo.Create(r.Context(), b); err != nil {
			httpx.Fail(w, http.StatusInternalServerError, "db", err.Error())
			return
		}
		httpx.JSON(w, http.StatusCreated, ToResponse(b))
	}
}

// @Summary List books
// @Tags    books
// @Produce json
// @Success 200 {object} httpx.Envelope{data=[]BookResponse}
// @Router  /v1/books [get]
func List(repo *Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		all, err := repo.GetAll(r.Context())
		if err != nil {
			httpx.Fail(w, http.StatusInternalServerError, "db", err.Error())
			return
		}
		out := make([]BookResponse, 0, len(all))
		for i := range all {
			out = append(out, ToResponse(&all[i]))
		}
		httpx.JSON(w, http.StatusOK, out)
	}
}

// @Summary Get book by id
// @Tags    books
// @Produce json
// @Param   id path int true "id"
// @Success 200 {object} httpx.Envelope{data=BookResponse}
// @Failure 404 {object} httpx.Envelope
// @Router  /v1/books/{id} [get]
func GetByID(repo *Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id64, _ := strconv.ParseUint(idStr, 10, 32)
		b, err := repo.GetByID(r.Context(), uint(id64))
		if err != nil {
			httpx.Fail(w, http.StatusNotFound, "not_found", "book not found")
			return
		}
		httpx.JSON(w, http.StatusOK, ToResponse(b))
	}
}

// @Summary Update book
// @Tags    books
// @Accept  json
// @Produce json
// @Param   id path int true "id"
// @Param   payload body UpdateBookRequest true "payload"
// @Success 200 {object} httpx.Envelope{data=BookResponse}
// @Failure 400 {object} httpx.Envelope
// @Failure 404 {object} httpx.Envelope
// @Router  /v1/books/{id} [put]
func Update(repo *Repo, v *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id64, _ := strconv.ParseUint(idStr, 10, 32)

		b, err := repo.GetByID(r.Context(), uint(id64))
		if err != nil {
			httpx.Fail(w, http.StatusNotFound, "not_found", "book not found")
			return
		}
		var in UpdateBookRequest
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			httpx.Fail(w, http.StatusBadRequest, "bad_json", "invalid json")
			return
		}
		if err := v.Struct(in); err != nil {
			httpx.Fail(w, http.StatusBadRequest, "validation", err.Error())
			return
		}

		if in.Name != nil {
			b.Name = *in.Name
		}
		if in.Author != nil {
			b.Author = *in.Author
		}
		if in.Publication != nil {
			b.Publication = *in.Publication
		}

		if err := repo.Save(r.Context(), b); err != nil {
			httpx.Fail(w, http.StatusInternalServerError, "db", err.Error())
			return
		}
		httpx.JSON(w, http.StatusOK, ToResponse(b))
	}
}

// @Summary Delete book
// @Tags    books
// @Produce json
// @Param   id path int true "id"
// @Success 204 {object} httpx.Envelope
// @Failure 404 {object} httpx.Envelope
// @Router  /v1/books/{id} [delete]
func Delete(repo *Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id64, _ := strconv.ParseUint(idStr, 10, 32)
		if err := repo.Delete(r.Context(), uint(id64)); err != nil {
			httpx.Fail(w, http.StatusNotFound, "not_found", "book not found")
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
