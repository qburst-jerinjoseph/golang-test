package app

import (
	"lazy-go/internal/core"
	"lazy-go/internal/data"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/sirupsen/logrus"
)

// Handler for API.
type Handler struct {
	logrus.FieldLogger
	data.Repo
}

// NewHandler API handler.
func NewHandler(logger logrus.FieldLogger, repo data.Repo) *Handler {
	return &Handler{logger, repo}
}

// Route v8 traffic.
func (h *Handler) Route() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.handlerExample)
	return r
}

func (h *Handler) handlerExample(w http.ResponseWriter, r *http.Request) {
	c := core.Context(r)
	data, err := h.GetSample(c)
	if err != nil {
		h.Error("Database operation failed", err)
		fail(w, http.StatusInternalServerError, 0, "Database operation failed.")
		return
	}
	send(w, http.StatusOK, data)
}
