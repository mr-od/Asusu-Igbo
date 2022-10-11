-- name: SearchProduct :many
SELECT *
FROM products
WHERE tsv @@ to_tsquery(sqlc.arg(search_query))
ORDER BY created_at DESC
LIMIT 100;