package api

import (
	"net/http"

	"github.com/justcompile/tfreg/api/handlers"
	"github.com/justcompile/tfreg/internal"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// NewRouter creates a top-level HTTP handler which contains url => handler or url => sub-router mappings
func NewRouter(app *internal.Application) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.RequestID,
		middleware.Logger,          // Log API request calls
		middleware.DefaultCompress, // Compress results, mostly gzipping assets and json
		middleware.Recoverer,       // Recover from panics without crashing server
		middleware.URLFormat,
		middleware.StripSlashes,
	)

	router.Get("/health", func(resp http.ResponseWriter, req *http.Request) {
		render.PlainText(resp, req, "Ok")
	})

	router.Route("/v1/modules", func(subRouter chi.Router) {
		subRouter.Get("/", handlers.Index)
	})

	return router
}
