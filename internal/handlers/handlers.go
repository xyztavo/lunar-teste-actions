package handlers

import (
	"bytes"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lunai-monster/lunar-pos/internal/config"
	"github.com/lunai-monster/lunar-pos/internal/database/sqlc"
	"github.com/lunai-monster/lunar-pos/internal/models"
	"github.com/lunai-monster/lunar-pos/internal/utils"
	"github.com/lunai-monster/lunar-pos/templates"
	"github.com/lunai-monster/lunar-pos/templates/fragments"
)

type Handler struct {
	cfg *config.Config
	q   *sqlc.Queries
	u   *utils.Utils
}

func NewHandlers(cfg *config.Config, q *sqlc.Queries, u *utils.Utils) *Handler {
	return &Handler{
		cfg: cfg,
		q:   q,
		u:   u,
	}
}

func (h *Handler) GetIndex(w http.ResponseWriter, r *http.Request) {
	products, err := h.q.GetProducts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := templates.Index(products).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := new(models.ProductRequest)
	if err := h.u.BindAndValidate(r, product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	arg := sqlc.InsertProductParams{
		Sku:        product.SKU,
		Title:      product.Title,
		Pricecents: int64(product.PriceCents),
	}
	_, err := h.q.InsertProduct(r.Context(), arg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products, err := h.q.GetProducts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := fragments.ProductList(products).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	sku := chi.URLParam(r, "sku")
	if err := h.q.DeleteProduct(r.Context(), sku); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products, err := h.q.GetProducts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := fragments.ProductList(products).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// update product
func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	sku := chi.URLParam(r, "sku")
	if sku == "" {
		http.Error(w, "missing sku", http.StatusBadRequest)
		return
	}

	product := new(models.ProductRequest)
	if err := h.u.BindAndValidate(r, product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	arg := sqlc.UpdateProductParams{
		Sku:        sku,
		Title:      product.Title,
		Pricecents: int64(product.PriceCents),
	}
	if err := h.q.UpdateProduct(r.Context(), arg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products, err := h.q.GetProducts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	listBuf := &bytes.Buffer{}
	if err := fragments.ProductList(products).Render(r.Context(), listBuf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	formBuf := &bytes.Buffer{}
	if err := templates.ProductForm(nil).Render(r.Context(), formBuf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(listBuf.Bytes())
	w.Write([]byte(`<div id="product-form">`))
	w.Write(formBuf.Bytes())
	w.Write([]byte(`</div>`))
}

func (h *Handler) GetProductForm(w http.ResponseWriter, r *http.Request) {
	sku := chi.URLParam(r, "sku")
	if sku == "" {
		http.Error(w, "missing sku", http.StatusBadRequest)
		return
	}

	product, err := h.q.GetProduct(r.Context(), sku)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	formBuf := &bytes.Buffer{}
	if err := templates.ProductForm(&product).Render(r.Context(), formBuf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`<div id="product-form">`))
	w.Write(formBuf.Bytes())
	w.Write([]byte(`</div>`))
}
