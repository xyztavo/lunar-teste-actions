package models

type ProductRequest struct {
	SKU        string `json:"sku" schema:"sku"`
	Title      string `json:"title" schema:"title"`
	PriceCents int    `json:"priceCents" schema:"priceCents"`
}
