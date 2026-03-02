-- name: InsertProduct :one
INSERT INTO products (sku, title, priceCents) VALUES (?, ?, ?) RETURNING sku;
-- name: GetProducts :many
SELECT * FROM products;

-- name: GetProduct :one
SELECT * FROM products WHERE sku = ?;

-- name: DeleteProduct :exec
DELETE FROM products WHERE sku = ?;

-- name: UpdateProduct :exec
UPDATE products SET title = ?, priceCents = ? WHERE sku = ?;