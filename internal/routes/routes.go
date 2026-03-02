package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lunai-monster/lunar-pos/internal/config"
	"github.com/lunai-monster/lunar-pos/internal/handlers"
)

func NewRouter(h *handlers.Handler, cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/*", http.StripPrefix("/", fs))
	r.Get("/", h.GetIndex)
	r.Post("/products", h.CreateProduct)
	r.Get("/products/{sku}/edit", h.GetProductForm)
	r.Delete("/products/{sku}", h.DeleteProduct)
	r.Put("/products/{sku}", h.UpdateProduct)
	return r
}
