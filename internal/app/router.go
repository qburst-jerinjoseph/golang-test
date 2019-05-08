package app

import (
	"lazy-go/internal/core/log"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Router sets up shohin'EtagHeader router
func (s *Server) Router() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.DefaultCompress)
	r.Use(middleware.RealIP)
	r.Use(log.NewStructuredLogger(s.Logger))
	r.Use(middleware.RequestID)
	r.Use(middleware.StripSlashes)
	h := NewHandler(s, s)
	r.Mount("/", h.Route())
	return r
}
